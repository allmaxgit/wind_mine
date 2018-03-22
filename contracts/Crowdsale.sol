pragma solidity ^0.4.18;

import './SafeMath.sol';
import './Ownable.sol';
import './WindMineToken.sol';
import './UsingFiatPrice.sol';

/**
 * @title Crowdsale
 * @dev WindMine token crowdsale. Extends UsingFiatPrice for updating price in fiat currency from external source
 */
contract Crowdsale is UsingFiatPrice {
    using SafeMath for uint256;

    enum State {
        NOT_STARTED,
        PRIVATE,
        PRE_ICO,
        ICO,
        FINISHED
    }

    /**
     * @dev Freezing period for frozen reserve of 80 mln tokens
     */
    uint256 public constant freezePeriod = 1 years;

    /**
     * @dev Amount of frozen tokens without token decimals
     */
    uint256 public constant frozenReserve = 80000000;

    /**
     * @dev Timestamp of started freezing date
     */
    uint256 public reserveFreezeTimestamp;

    /**
     * @title investorsList
     * @dev List of investors to get investors statistics
     */
    address[] public investorsList;

    /**
     * @title privatePriceInFiatFracture
     * @dev Token price in minimal fiat currency fracture for Private Sale stage
     */
    uint256 public privatePriceInFiatFracture;

    /**
     * @title preIcoPriceInFiatFracture
     * @dev Token price in minimal fiat currency fracture for Pre-ICO stage
     */
    uint256 public preIcoPriceInFiatFracture;

    /**
     * @title icoPriceInFiatFracture
     * @dev Token price in minimal fiat currency fracture for ICO stage
     */
    uint256 public icoPriceInFiatFracture;

    /**
     * @title generalHardCap
     * @dev General hard cap, which is defined by token total supply
     */
    uint256 public generalHardCap;

    /**
     * @title currentHardCap
     * @dev Hard cap for current crowdsale stage
     */
    uint256 public currentHardCap;

    /**
     * @dev privateSaleHardCap
     * @title Hard cap for Private Sale stage
     */
    uint256 public privateSaleHardCap;

    /**
     * @title preIcoHardCap
     * @dev Hard cap for Pre-ICO stage, which includes previous stage hard cap
     * @notice This hard cap includes previous stage hard cap, so it is relative to general amount of tokens
     */
    uint256 public preIcoHardCap;

    /**
     * @title icoHardCap
     * @dev Hard cap for ICO stage, which includes previous stage hard cap
     * @notice This hard cap includes previous stage hard cap, so it is relative to general amount of tokens
     */
    uint256 public icoHardCap;

    /**
     * @title privateSaleStartDate
     * @dev Private Sale stage start date
     */
    uint256 public privateSaleStartDate;

    /**
     * @title privateSaleDuration
     * @dev Duration of the Private Sale stage
     */
    uint256 public privateSaleDuration;

    /**
     * @title preIcoStartDate
     * @dev Pre-ICO stage start date
     * @notice It is calculated by summing Private Sale stage start time and Private Sale stage duration
     */
    uint256 public preIcoStartDate;

    /**
     * @title preIcoDuration
     * @dev Duration of the Pre-ICO stage
     */
    uint256 public preIcoDuration;

    /**
     * @title icoStartDate
     * @dev ICO stage start date
     * @notice It is calculated by summing Pre-ICO stage start time and Pre-ICO stage duration
     */
    uint256 public icoStartDate;

    /**
     * @title icoDuration
     * @dev Duration of the ICO stage
     */
    uint256 public icoDuration;

    /**
     * @title tokensSold
     * @dev Number of sold tokens
     */
    uint256 public tokensSold;

    /**
     * @title weiRaised
     * @dev Amount of raised funds in wei
     */
    uint256 public weiRaised;

    /**
     * @title wallet
     * @dev Beneficiary wallet address
     */
    address public wallet;

    /**
     * @title crowdsaleState
     * @dev Current crowdsale state
     * @notice This state is updated in every transaction by calling checkState function
     */
    State public crowdsaleState = State.NOT_STARTED;

    /**
     * @title token
     * @dev WindMineToken
     */
    WindMineToken public token;

    /**
     * @dev Mapping of number of tokens ordered during sale stages
     */
    mapping (address => uint256) public tokensOrdered;

    mapping (address => bool) public hasPassedKYC;

    /**
     * @dev Mapping of private participants
     * @notice If bool flag is set to true, then this address is allowed to participate in the Private Sale stage
     */
    mapping (address => bool) private privateParticipants;

    /**
     * @dev We use a single lock for the whole contract.
     */
    bool private reentrancy_lock = false;

    /**
     * @dev Prevents a contract from calling itself, directly or indirectly.
     */
    modifier nonReentrant() {
        require(!reentrancy_lock);
        reentrancy_lock = true;
        _;
        reentrancy_lock = false;
    }

    /**
     * @dev This event is raised when tokens were manually send
     * @notice Backend uses manual sending for received bitcoin transactions from registered users
     */
    event ManualTokenSend(address _receiver, uint256 _amount);

    /**
     * @dev This event is raised when crowdsale state changes
     */
    event StateHasChanged(State _oldState, State _newState);

    /**
     * @dev This event is raised when owner withdraws raised funds
     */
    event FundsWithdrawn(address _wallet, uint256 _amount);

    /**
     * @dev This event is raised when owner changes beneficiary wallet
     */
    event WalletHasChanged(address _oldWallet, address _newWallet);

    /**
     * @dev This event is raised when owner moves private sale start date (and all other start dates with it)
     */
    event StartDateMoved(uint256 _oldDate, uint256 _newDate);

    /**
     * @dev This event is raised when investor order tokens during any sale period by transferring ETH to this contract
     */
    event TokensAreOrdered(address _orderer, uint256 _amount, uint256 _currentWeiInFiat);

    /**
     * @dev This event is raised when investor retrieves his ordered tokens after ICO end
     */
    event TokensAreRetrieved(address _retriever, uint256 _amount);

    /**
     * @dev This event is raised when owner retrieves frozen reserve of 80 mln tokens after 12 month
     */
    event FrozenReserveRetrieved(address _receiver, uint256 _unfreezeTimestamp);

    /**
     * @dev Crowdsale constructor, which calls parent UsingFiatPrice constructor
     * @param _privateSaleStartDate Private sale stage start date in Unix timestamp
     * @param _privateSaleDuration Private sale stage duration in Unix timestamp
     * @param _preIcoDuration Pre-ICO stage duration in Unix timestamp
     * @param _icoDuration ICO stage duration in Unix timestamp
     * @param _wallet Wallet where raised funds will be sent when crowdsale finishes
     * @param _foundersWallet Founder's wallet address, where to send 20 mln tokens of founder's reserve
     * @notice Crowdsale constructor calls parent contract UsingFiatPrice constructor, setting fiat symbol to EUR and
     * fiat decimals to 2
     */
    function Crowdsale(uint256 _privateSaleStartDate, uint256 _privateSaleDuration, uint256 _preIcoDuration, uint256 _icoDuration, address _wallet, address _foundersWallet)
    UsingFiatPrice("EUR", 2) public {
        require(_privateSaleStartDate > 0);
        require(_privateSaleDuration > 0);
        require(_preIcoDuration > 0);
        require(_icoDuration > 0);
        require(_wallet != address(0));
        require(_foundersWallet != address(0));

        setDates(_privateSaleStartDate, _privateSaleDuration, _preIcoDuration, _icoDuration);
        wallet = _wallet;
        token = new WindMineToken(_foundersWallet);

        reserveFreezeTimestamp = now;
    }

    function getInvestorsListLength() public view returns (uint256) {
        return investorsList.length;
    }

    /**
     * @dev Users can send ETH directly to the crowdsale to buy tokens
     */
    function() public payable {
        buyTokens(msg.sender);
    }

    /**
     * @dev Function for handling received ETH and sending tokens to sender
     * @param _sender Address of funds sender
     */
    function buyTokens(address _sender) public payable nonReentrant {
        checkState();
        require(_sender != address(0));
        require(crowdsaleState != State.NOT_STARTED);
        require(crowdsaleState != State.FINISHED);
        require(msg.value > 0);
        require(weiInFiat > 0);

        if (crowdsaleState == State.PRIVATE) {
            require(privateParticipants[_sender]);
        }
        //KYC check
        require(hasPassedKYC[_sender]);

        uint256 receivedWei = msg.value;
        uint256 fiatAmount = receivedWei.div(weiInFiat);
        //Fiat equivalent of received wei is bigger than one integral unit of fiat currency
        require(fiatAmount >= 10 ** fiatDecimals);
        uint256 tokensBought;
        uint256 stagePriceInFiatFracture;
        uint256 change;
        uint256 realReceivedWei;
        if (crowdsaleState == State.PRIVATE) {
            stagePriceInFiatFracture = privatePriceInFiatFracture;
        } else if (crowdsaleState == State.PRE_ICO) {
            stagePriceInFiatFracture = preIcoPriceInFiatFracture;
        } else if (crowdsaleState == State.ICO) {
            stagePriceInFiatFracture = icoPriceInFiatFracture;
        }
        tokensBought = fiatAmount.div(stagePriceInFiatFracture);

        require(tokensBought > 0);

        if (tokensSold.add(tokensBought) > currentHardCap) {
            tokensBought = currentHardCap.sub(tokensSold);
            realReceivedWei = tokensBought.mul(stagePriceInFiatFracture).mul(weiInFiat);
            change = receivedWei.sub(realReceivedWei);
        } else {
            realReceivedWei = receivedWei;
        }

        tokensSold = tokensSold.add(tokensBought);
        weiRaised = weiRaised.add(realReceivedWei);

        if (token.balanceOf(_sender) == 0) {
            investorsList.push(_sender);
        }

        if (change > 0) {
            _sender.transfer(change);
        }

        tokensBought = tokensBought.mul(10 ** token.decimals());
        tokensOrdered[_sender] = tokensOrdered[_sender].add(tokensBought);

        TokensAreOrdered(_sender, tokensBought, weiInFiat);
    }

    /**
     * @dev Sends tokens, ordered during sale periods, to the caller
     */
    function receiveOrderedTokens() public nonReentrant {
        require(crowdsaleState == State.FINISHED);
        require(tokensOrdered[msg.sender] > 0);
        //KYC check
        require(hasPassedKYC[msg.sender]);

        uint256 numOfTokens = tokensOrdered[msg.sender];
        tokensOrdered[msg.sender] = 0;

        token.transfer(msg.sender, numOfTokens);
        TokensAreRetrieved(msg.sender, numOfTokens);
    }

    /**
     * @dev Checks crowdsale state and updates state field and current stage hard cap
     */
    function checkState() public {
        if (now < privateSaleStartDate) {
            //Crowdsale is not started yet
            crowdsaleState = State.NOT_STARTED;
        } else if (now >= privateSaleStartDate && now < preIcoStartDate) {
            //Crowdsale is in Private Sale state. Set state to PRIVATE and update hard cap to private sale hard cap
            StateHasChanged(State.NOT_STARTED, State.PRIVATE);
            crowdsaleState = State.PRIVATE;
            currentHardCap = privateSaleHardCap;
        } else if (now >= preIcoStartDate && now < icoStartDate) {
            //Crowdsale is in Pre-ICO state. Set state to PRE_ICO and update hard cap to Pre-ICO hard cap
            StateHasChanged(State.PRIVATE, State.PRE_ICO);
            crowdsaleState = State.PRE_ICO;
            currentHardCap = preIcoHardCap;
        } else if (now >= icoStartDate && now < icoStartDate + icoDuration) {
            //Crowdsale is in ICO state. Set state to ICO and update hard cap to ICO hard cap
            StateHasChanged(State.PRE_ICO, State.ICO);
            crowdsaleState = State.ICO;
            currentHardCap = icoHardCap;
        } else {
            //Crowdsale has finished. Set state to FINISHED
            StateHasChanged(State.ICO, State.FINISHED);
            crowdsaleState = State.FINISHED;
        }
    }

    /**
     * @dev Mark address as KYC verified
     * @param _investor Address of the investor, who passed KYC
     * @notice After approval this investor will be able to buy tokens
     */
    function approveKYC(address _investor) public onlyOwner {
        require(_investor != address(0));
        hasPassedKYC[_investor] = true;
    }

    /**
     * @dev Mark list of addresses as KYC verified
     * @param _investors List of address of the investors, who has passed KYC
     * @notice After approval this investors will be able to buy tokens
     * @notice Cycles tend to deplete all available gas, so array length bound is introduced
     */
    function approveKYCForList(address[] _investors) public onlyOwner {
        uint256 length = _investors.length;
        require(length > 0 && length <= 150); //TODO experiment with higher bound

        for (uint i = 0; i < length; i++) {
            approveKYC(_investors[i]);
        }
    }

    /**
     * @dev Mark address as KYC non-verified
     * @param _investor Address of the investor, who has failed KYC
     * @notice After disapproval this investor will not be able to buy tokens
     */
    function disapproveKYC(address _investor) public onlyOwner {
        require(_investor != address(0));
        hasPassedKYC[_investor] = false;
    }

    /**
     * @dev Mark list of addresses as KYC non-verified
     * @param _investors List of address of the investors, who has failed KYC
     * @notice After disapproval this investor will not be able to buy tokens
     * @notice Cycles tend to deplete all available gas, so array length bound is introduced
     */
    function disapproveKYCForList(address[] _investors) public onlyOwner {
        uint256 length = _investors.length;
        require(length > 0 && length <= 150); //TODO experiment with higher bound

        for (uint i = 0; i < length; i++) {
            disapproveKYC(_investors[i]);
        }
    }

    /**
     * @dev Sends frozen reserve to receiver if 12 month has passed since freezing
     * @param _receiver Address of the frozen reserve receiver
     */
    function retrieveFrozenReserve(address _receiver) public onlyOwner nonReentrant {
        require(now >= reserveFreezeTimestamp + freezePeriod);
        require(_receiver != address(0));

        token.transfer(_receiver, frozenReserve.mul(10 ** token.decimals()));
        FrozenReserveRetrieved(_receiver, now);
    }

    /**
     * @dev Prepares crowdsale by setting hard caps for stages using information from token
     */
    function prepareCrowdsale() public onlyOwner {
        generalHardCap = token.totalSupply();
        privateSaleHardCap = 10000000 * (10 ** token.decimals());
        preIcoHardCap = 20000000 * (10 ** token.decimals()) + privateSaleHardCap;
        icoHardCap = 30000000 * (10 ** token.decimals()) + preIcoHardCap;
        privatePriceInFiatFracture = (10 ** fiatDecimals).mul(25).div(100);
        preIcoPriceInFiatFracture = (10 ** fiatDecimals).mul(50).div(100);
        icoPriceInFiatFracture = 10 ** fiatDecimals;
        checkState();
    }

    /**
     * @dev Allows this address to participate in Private Sale stage
     * @param _participant Wallet address of the private participant
     */
    function addPrivateParticipant(address _participant) public onlyOwner {
        require(_participant != address(0));
        privateParticipants[_participant] = true;
    }

    /**
     * @dev Sets new wallet where raised funds should be sent after the end of crowdsale
     * @param _newWallet New wallet address
     * @notice It is possible to change wallet only if crowdsale isn't finished yet
     */
    function setWallet(address _newWallet) public onlyOwner {
        checkState();
        require(crowdsaleState != State.FINISHED);
        WalletHasChanged(wallet, _newWallet);
        wallet = _newWallet;
    }

    /**
     * @dev Sets new starting date for the first stage - private sale, moving all other dates by saved stage durations
     * @param _newDate New Private Sale stage starting date
     * @notice It is possible to change starting date only if crowdsale is not started yet
     */
    function setNewStartDate(uint256 _newDate) public onlyOwner {
        checkState();
        require(crowdsaleState == State.NOT_STARTED);
        StartDateMoved(privateSaleStartDate, _newDate);
        setDates(_newDate, privateSaleDuration, preIcoDuration, icoDuration);
    }

    /**
     * @dev Manually reserves tokens for receiver
     * @param _receiver Token receiver wallet address
     * @param _amount Amount of tokens to send
     * @notice Intended to be used by backend to manually give tokens to BTC investors
     */
    function manualReserve(address _receiver, uint256 _amount) public onlyOwner nonReentrant {
        checkState();
        require(_receiver != address(0));
        require(_amount > 0);
        require(crowdsaleState != State.NOT_STARTED && crowdsaleState != State.FINISHED);
        require(tokensSold.add(_amount) < currentHardCap);

        if (crowdsaleState == State.PRIVATE) {
            require(privateParticipants[_receiver]);
        }

        //KYC check
        require(hasPassedKYC[_receiver]);

        if (token.balanceOf(_receiver) == 0) {
            investorsList.push(_receiver);
        }

        tokensSold = tokensSold.add(_amount);
        tokensOrdered[_receiver] = tokensOrdered[_receiver].add(_amount);

        TokensAreOrdered(_receiver, _amount, weiInFiat);
    }

    /**
     * @dev Sends raised funds to the beneficiary wallet
     * @notice It is possible to withdraw raised funds only if crowdsale is finished
     */
    function withdraw() public onlyOwner nonReentrant {
        checkState();
        require(crowdsaleState == State.FINISHED);

        FundsWithdrawn(wallet, weiRaised);
        wallet.transfer(weiRaised);
    }

    /**
     * @dev Set stages start dates and durations
     * @param _privateSaleStartDate Private Sale stage start date in Unix timestamp
     * @param _privateSaleDuration Private Sale stage duration in Unix timestamp
     * @param _preIcoDuration Pre-ICO stage duration in Unix timestamp
     * @param _icoDuration ICO stage duration in Unix timestamp
     */
    function setDates(uint256 _privateSaleStartDate, uint256 _privateSaleDuration, uint256 _preIcoDuration, uint256 _icoDuration) private {
        privateSaleStartDate = _privateSaleStartDate;
        privateSaleDuration = _privateSaleDuration;
        preIcoStartDate = privateSaleStartDate + privateSaleDuration;
        preIcoDuration = _preIcoDuration;
        icoStartDate = preIcoStartDate + preIcoDuration;
        icoDuration = _icoDuration;
    }
}
