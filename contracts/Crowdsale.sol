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
     * @title freezePeriod
     * @dev Freezing period for frozen reserve of 80 mln tokens
     */
    uint256 public constant freezePeriod = 1 years;

    /**
     * @title frozenReserve
     * @dev Amount of frozen tokens without token decimals
     */
    uint256 public constant frozenReserve = 80000000;

    /**
     * @title isReserveSent
     * @dev Flag which shows, whether frozen reserve has been sent or not
     */
    bool public isReserveSent = false;

    /**
     * @title reserveFreezeTimestamp
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
     * @title Hard cap for Private Sale stage
     * @dev privateSaleHardCap
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
     * @title preIcoStartDate
     * @dev Pre-ICO stage start date
     */
    uint256 public preIcoStartDate;

    /**
     * @title icoStartDate
     * @dev ICO stage start date
     */
    uint256 public icoStartDate;

    /**
     * @title icoFinishDate
     * @dev ICO stage finish date
     */
    uint256 public icoFinishDate;

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
     * @title privateSaleWeiRaised
     * @dev Amount of raised funds in wei at Private Sale stage
     */
    uint256 public privateSaleWeiRaised;

    /**
     * @title privateSaleFundsWithdrawn
     * @dev Flags if funds raised at Private Sale has been withdrawn
     */
    bool public privateSaleFundsWithdrawn = false;

    /**
     * @title preIcoWeiRaised
     * @dev Amount of raised funds in wei at Pre-ICO stage
     */
    uint256 public preIcoWeiRaised;

    /**
     * @title preIcoFundsWithdrawn
     * @dev Flags if funds raised at Pre-ICO has been withdrawn
     */
    bool public preIcoFundsWithdrawn = false;

    /**
     * @title icoWeiRaised
     * @dev Amount of raised funds in wei at ICO stage
     */
    uint256 public icoWeiRaised;

    /**
     * @title icoFundsWithdrawn
     * @dev Flags if funds raised at ICO has been withdrawn
     */
    bool public icoFundsWithdrawn = false;

    bool[4] public withdrawnState;

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
     * @title tokensOrdered
     * @dev Mapping of number of tokens ordered during sale stages
     */
    mapping (address => uint256) public tokensOrdered;

    /**
     * @title totalWhiteListed
     * @dev Total amount of whitelisted addresses
     */
    uint256 public totalWhiteListed;

    /**
     * @title whiteList
     * @dev A mapping of addresses to flags, which shows whether address has passed KYC or not
     */
    mapping(address => bool) public whiteList;

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
     * @param _receiver The receiver of the manually reserved tokens
     * @param _amount The amount of tokens to be reserved for user
     * @notice Amount of tokens should be with token decimals
     * @notice Backend uses manual sending for received bitcoin transactions from registered users
     */
    event ManualTokenSend(address _receiver, uint256 _amount);

    /**
     * @dev This event is raised when crowdsale state changes
     * @param _oldState Old crowdsale state
     * @param _newState New crowdsale state
     */
    event StateHasChanged(State _oldState, State _newState);

    /**
     * @dev This event is raised when owner withdraws raised funds
     * @param _wallet Where withdrawn funds were sent
     * @param _amount How much wei were sent
     */
    event FundsWithdrawn(address _wallet, uint256 _amount, uint256 _stage);

    /**
     * @dev This event is raised when owner changes beneficiary wallet
     * @param _oldWallet Old beneficiary wallet
     * @param _newWallet New beneficiary wallet
     */
    event WalletHasChanged(address _oldWallet, address _newWallet);

    /**
     * @dev This event is raised when owner moves Private Sale start date
     * @param _oldDate Old start date of Private Sale
     * @param _newDate New start date of Private Sale
     */
    event PrivateSaleStartDateMoved(uint256 _oldDate, uint256 _newDate);

    /**
     * @dev This event is raised when owner moves Pre-ICO start date
     * @param _oldDate Old start date of Pre-ICO
     * @param _newDate New start date of Pre-ICO
     */
    event PreIcoStartDateMoved(uint256 _oldDate, uint256 _newDate);

    /**
     * @dev This event is raised when owner moves ICO start date
     * @param _oldDate Old start date of ICO
     * @param _newDate New start date of ICO
     */
    event IcoStartDateMoved(uint256 _oldDate, uint256 _newDate);

    /**
     * @dev This event is raised when owner moves ICO finish date
     * @param _oldDate Old finish date of ICO
     * @param _newDate New finish date of ICO
     */
    event IcoFinishDateMoved(uint256 _oldDate, uint256 _newDate);

    /**
     * @dev This event is raised when investor order tokens during any sale period by transferring ETH to this contract
     * @param _orderer Address of ETH sender
     * @param _amount How much tokens ordered (with decimals)
     * @param _currentWeiInFiat Current amount of wei in 0.01 EUR
     */
    event TokensAreOrdered(address _orderer, uint256 _amount, uint256 _currentWeiInFiat);

    /**
     * @dev This event is raised when investor retrieves his ordered tokens after ICO end
     * @param _retriever Who has retrieved his ordered tokens
     * @param _amount How much tokens he was sent (with decimals)
     */
    event TokensAreRetrieved(address _retriever, uint256 _amount);

    /**
     * @dev This event is raised when owner retrieves frozen reserve of 80 mln tokens after 12 month
     * @param _receiver Receiver of frozen token reserve
     * @param _unfreezeTimestamp Timestamp of retrieval date
     */
    event FrozenReserveRetrieved(address _receiver, uint256 _unfreezeTimestamp);

    /**
     * @dev This event is raised, when address `user` has been whitelisted
     * @param user Who was whitelisted
     * @param whiteListedNum How much addresses are currently whitelisted (after adding)
     */
    event LogWhiteListed(address user, uint256 whiteListedNum);

    /**
     * @dev This event is raised, when owner whitelists an array of addresses
     * @param whiteListedNum How much addresses are currently whitelisted (after adding)
     */
    event LogWhiteListedMultiple(uint256 whiteListedNum);

    /**
     * @dev Crowdsale constructor, which calls parent UsingFiatPrice constructor
     * @param _privateSaleStartDate Private sale stage start date in Unix timestamp
     * @param _preIcoStartDate Private sale stage start date in Unix timestamp
     * @param _icoStartDate Private sale stage start date in Unix timestamp
     * @param _icoFinishDate Private sale stage finish date in Unix timestamp
     * @param _wallet Wallet where raised funds will be sent when crowdsale finishes
     * @param _foundersWallet Founder's wallet address, where to send 20 mln tokens of founder's reserve
     * @notice Crowdsale constructor calls parent contract UsingFiatPrice constructor, setting fiat symbol to EUR and
     * fiat decimals to 2
     */
    function Crowdsale(uint256 _privateSaleStartDate, uint256 _preIcoStartDate, uint256 _icoStartDate, uint256 _icoFinishDate, address _wallet, address _foundersWallet)
    UsingFiatPrice("EUR", 2) public {
        require(_privateSaleStartDate > 0);
        require(_preIcoStartDate > _privateSaleStartDate);
        require(_icoStartDate > _preIcoStartDate);
        require(_wallet != address(0));
        require(_foundersWallet != address(0));

        privateSaleStartDate = _privateSaleStartDate;
        preIcoStartDate = _preIcoStartDate;
        icoStartDate = _icoStartDate;
        icoFinishDate = _icoFinishDate;
        wallet = _wallet;
        token = new WindMineToken(_foundersWallet);
    }

    /**
     * @dev Gets number of investors
     * @return number of investors
     */
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
        require(msg.value >= 1E17); //Minimal purchase is 0.1 ETH
        require(weiInFiat > 0);

        if (crowdsaleState == State.PRIVATE) {
            require(privateParticipants[_sender]);
        }
        //KYC check
        require(whiteList[_sender]);

        uint256 receivedWei = msg.value;
        uint256 fiatAmount = receivedWei.div(weiInFiat);
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
        tokensBought = fiatAmount.mul(10 ** token.decimals()).div(stagePriceInFiatFracture); //amount of tokens bought with decimals

        require(tokensBought > 0);
        if (tokensSold.add(tokensBought) > currentHardCap) {
            tokensBought = currentHardCap.sub(tokensSold);
            realReceivedWei = tokensBought.mul(stagePriceInFiatFracture).mul(weiInFiat).div(10 ** token.decimals());
            change = receivedWei.sub(realReceivedWei);
        } else {
            realReceivedWei = receivedWei;
        }

        tokensSold = tokensSold.add(tokensBought);
        weiRaised = weiRaised.add(realReceivedWei);
        if (crowdsaleState == State.PRIVATE) {
            privateSaleWeiRaised = privateSaleWeiRaised.add(realReceivedWei);
        } else if (crowdsaleState == State.PRE_ICO) {
            preIcoWeiRaised = preIcoWeiRaised.add(realReceivedWei);
        } else if (crowdsaleState == State.ICO) {
            icoWeiRaised = icoWeiRaised.add(realReceivedWei);
        }

        if (tokensOrdered[_sender] == 0) {
            investorsList.push(_sender);
        }

        if (change > 0) {
            _sender.transfer(change);
        }

        tokensOrdered[_sender] = tokensOrdered[_sender].add(tokensBought);

        TokensAreOrdered(_sender, tokensBought, weiInFiat);
    }

    /**
     * @dev Sends tokens, ordered during sale periods, to the caller
     */
    function receiveOrderedTokens() public nonReentrant {
        require(crowdsaleState == State.FINISHED);
        require(tokensOrdered[msg.sender] > 0);

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
            if (crowdsaleState != State.PRIVATE) {
                StateHasChanged(State.NOT_STARTED, State.PRIVATE);
                crowdsaleState = State.PRIVATE;
                currentHardCap = privateSaleHardCap;
            }
        } else if (now >= preIcoStartDate && now < icoStartDate) {
            //Crowdsale is in Pre-ICO state. Set state to PRE_ICO and update hard cap to Pre-ICO hard cap
            if (crowdsaleState != State.PRE_ICO) {
                StateHasChanged(State.PRIVATE, State.PRE_ICO);
                crowdsaleState = State.PRE_ICO;
                currentHardCap = preIcoHardCap;
            }
        } else if (now >= icoStartDate && now < icoFinishDate) {
            //Crowdsale is in ICO state. Set state to ICO and update hard cap to ICO hard cap
            if (crowdsaleState != State.ICO) {
                StateHasChanged(State.PRE_ICO, State.ICO);
                crowdsaleState = State.ICO;
                currentHardCap = icoHardCap;
            }
        } else {
            //Crowdsale has finished. Set state to FINISHED
            if (crowdsaleState != State.FINISHED) {
                StateHasChanged(State.ICO, State.FINISHED);
                crowdsaleState = State.FINISHED;
                reserveFreezeTimestamp = icoFinishDate;
            }
        }
    }

    /**
     * @dev Whitelist address for token purchase
     * @param _user Address of the investor, who should be whitelisted
     * @notice After this address is added, it will be able to buy tokens (except for Private Sale)
     */
    function addToWhiteList(address _user) external onlyOwner returns (bool) {
        require(_user != address(0));
        if (!whiteList[_user]) {
            whiteList[_user] = true;
            totalWhiteListed++;
            LogWhiteListed(_user, totalWhiteListed);
        }
        return true;
    }

    /**
     * @dev Mark list of addresses as whitelisted
     * @param _users List of address of the investors, who should be whitelisted
     * @notice After this investors are added, they will be able to buy tokens
     * @notice Cycles tend to deplete all available gas, so array length constraint is introduced
     */
    function addToWhiteListMultiple(address[] _users) external onlyOwner returns (bool) {
        uint256 length = _users.length;
        require(length <= 150);
        for (uint i = 0; i < length; ++i) {
            if (!whiteList[_users[i]]) {
                whiteList[_users[i]] = true;
                totalWhiteListed++;
            }
        }
        LogWhiteListedMultiple(totalWhiteListed);
        return true;
    }

    /**
     * @dev Sends frozen reserve to receiver if 12 month has passed since freezing
     * @param _receiver Address of the frozen reserve receiver
     */
    function receiveFrozenReserve(address _receiver) public onlyOwner nonReentrant {
        require(crowdsaleState == State.FINISHED);
        require(now >= reserveFreezeTimestamp + freezePeriod);
        require(!isReserveSent);
        require(_receiver != address(0));

        isReserveSent = true;
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
        require(_newWallet != address(0));
        WalletHasChanged(wallet, _newWallet);
        wallet = _newWallet;
    }

    /**
     * @dev Sets new starting date of the Private Sale
     * @param _newDate New Private Sale stage starting date in Unix timestamp
     * @notice It is possible to change starting date only if crowdsale is not started yet
     * @notice New start date of Private Sale cannot be bigger than Pre-ICO start date
     */
    function setPrivateSaleStartDate(uint256 _newDate) public onlyOwner {
        checkState();
        require(crowdsaleState == State.NOT_STARTED);
        require(_newDate < preIcoStartDate); // private sale start date cannot be bigger than pre ico start date

        PrivateSaleStartDateMoved(privateSaleStartDate, _newDate);
        privateSaleStartDate = _newDate;
    }

    /**
     * @dev Sets new starting date of the Pre-ICO
     * @param _newDate New Pre-ICO stage starting date in Unix timestamp
     * @notice New start date of Pre-ICO cannot be bigger than ICO start date
     */
    function setPreIcoStartDate(uint256 _newDate) public onlyOwner {
        checkState();
        require(crowdsaleState == State.NOT_STARTED || crowdsaleState == State.PRIVATE);
        require(_newDate > privateSaleStartDate);
        require(_newDate < icoStartDate); // pre ico start date cannot be bigger than ico start date

        PreIcoStartDateMoved(preIcoStartDate, _newDate);
        preIcoStartDate = _newDate;
    }

    /**
     * @dev Sets new starting date of the ICO
     * @param _newDate New ICO stage starting date in Unix timestamp
     * @notice New start date of ICO cannot be bigger than ICO finish date
     */
    function setIcoStartDate(uint256 _newDate) public onlyOwner {
        checkState();
        require(crowdsaleState != State.ICO && crowdsaleState != State.FINISHED);
        require(_newDate > preIcoStartDate);
        require(_newDate < icoFinishDate); // pre ico start date cannot be bigger than ico finish date

        IcoStartDateMoved(icoStartDate, _newDate);
        icoStartDate = _newDate;
    }

    /**
     * @dev Sets new ending date of the ICO
     * @param _newDate New ICO stage starting date in Unix timestamp
     */
    function setIcoFinishDate(uint256 _newDate) public onlyOwner {
        checkState();
        require(crowdsaleState != State.FINISHED);
        require(_newDate > icoStartDate);

        IcoFinishDateMoved(icoFinishDate, _newDate);
        icoFinishDate = _newDate;
    }

    /**
     * @dev Manually reserves tokens for receiver
     * @param _receiver Token receiver wallet address
     * @param _amount Amount of tokens to send
     * @notice Intended to be used by backend to manually give tokens to BTC investors
     */
    function manualReserve(address _receiver, uint256 _amount) external onlyOwner nonReentrant {
        checkState();
        require(_receiver != address(0));
        require(_amount > 0);
        require(crowdsaleState != State.NOT_STARTED && crowdsaleState != State.FINISHED);
        require(tokensSold.add(_amount) < currentHardCap);

        if (crowdsaleState == State.PRIVATE) {
            require(privateParticipants[_receiver]);
        }

        //KYC check
        require(whiteList[_receiver]);

        if (token.balanceOf(_receiver) == 0) {
            investorsList.push(_receiver);
        }

        tokensSold = tokensSold.add(_amount);
        tokensOrdered[_receiver] = tokensOrdered[_receiver].add(_amount);

        TokensAreOrdered(_receiver, _amount, weiInFiat);
    }

    /**
     * @dev Sends raised funds to the beneficiary wallet
     * @notice It is possible to withdraw funds from stage only at the following stage. E.g. funds raised at private sale
     * can be withdrawn
     */
    function withdraw() public onlyOwner nonReentrant {
        checkState();
        require(crowdsaleState > State.NOT_STARTED);
        require(!withdrawnState[uint(crowdsaleState).sub(1)]);
        // Private Sale funds can be withdrawn only when state is Pre-ICO, ICO or Finished and funds haven't been withdrawn before
        if (crowdsaleState > State.PRIVATE && !withdrawnState[uint(State.PRIVATE)]) {
            withdrawnState[uint(State.PRIVATE)] = true;
            FundsWithdrawn(wallet, privateSaleWeiRaised, uint(State.PRIVATE));
            wallet.transfer(privateSaleWeiRaised);
        }
        // Pre-ICO funds can be withdrawn only when state is ICO or Finished and funds haven't been withdrawn before
        if (crowdsaleState > State.PRE_ICO && !withdrawnState[uint(State.PRE_ICO)]) {
            withdrawnState[uint(State.PRE_ICO)] = true;
            FundsWithdrawn(wallet, preIcoWeiRaised, uint(State.PRE_ICO));
            wallet.transfer(preIcoWeiRaised);
        }
        // ICO funds can be withdrawn only when state is Finished and funds haven't been withdrawn before
        if (crowdsaleState > State.ICO && !withdrawnState[uint(State.ICO)]) {
            withdrawnState[uint(State.ICO)] = true;
            FundsWithdrawn(wallet, icoWeiRaised, uint(State.ICO));
            wallet.transfer(icoWeiRaised);
        }
    }
}
