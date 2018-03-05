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
     * @notice This state is updated in every transaction by calling private checkState function
     */
    State public crowdsaleState = State.NOT_STARTED;

    /**
     * @title token
     * @dev WindMineToken
     */
    WindMineToken public token;

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

    event ManualTokenSend(address _receiver, uint256 _amount);
    event StateHasChanged(State _oldState, State _newState);
    event FundsWithdrawn(address _wallet, uint256 _amount);
    event WalletHasChanged(address _oldWallet, address _newWallet);
    event StartDateMoved(uint256 _oldDate, uint256 _newDate);

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

        ManualTokenSend(_receiver, _amount);

        if (token.balanceOf(_receiver) == 0) {
            investorsList.push(_receiver);
        }

        tokensSold = tokensSold.add(_amount);
        token.transfer(_receiver, _amount);
    }

    /**
     * @dev Sends raised funds to the beneficiary wallet
     * @notice It is possible to withdraw raised funds only if crowdsale is finished and the crowdsale target is reached
     */
    function withdraw() public onlyOwner nonReentrant {
        checkState();
        require(crowdsaleState == State.FINISHED);

        FundsWithdrawn(wallet, weiRaised);
        wallet.transfer(weiRaised);
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

        if (crowdsaleState == State.PRIVATE) {
            require(privateParticipants[_sender]);
        }

        uint256 sentWei = msg.value;
        uint256 fiatAmount = sentWei.div(weiInFiat);
        //Fiat equivalent of received wei is bigger than one integral unit of fiat currency
        require(fiatAmount >= 10 ** fiatDecimals);
        uint256 tokens;
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
        tokens = fiatAmount.div(stagePriceInFiatFracture);
        assert(tokens > 0);
        if (tokensSold.add(tokens) > currentHardCap) {
            tokens = currentHardCap.sub(tokensSold);
            realReceivedWei = tokens.mul(stagePriceInFiatFracture).mul(weiInFiat);
            change = sentWei.sub(realReceivedWei);
        } else {
            realReceivedWei = sentWei;
        }

        tokensSold = tokensSold.add(tokens);
        weiRaised = weiRaised.add(realReceivedWei);

        if (token.balanceOf(_sender) == 0) {
            investorsList.push(_sender);
        }

        if (change > 0) {
            _sender.transfer(change);
        }

        tokens = tokens.mul(10 ** token.decimals);
        token.transfer(_sender, tokens);
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
