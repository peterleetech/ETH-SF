/**
 *Submitted for verification at polygonscan.com on 2022-11-06
*/

// Sources flattened with hardhat v2.9.6 https://hardhat.org

// File @openzeppelin/contracts/utils/Context.sol@v4.6.0
// SPDX-License-Identifier: MIT

// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

pragma solidity ^0.8.0;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}


// File @openzeppelin/contracts/access/Ownable.sol@v4.6.0


// OpenZeppelin Contracts v4.4.1 (access/Ownable.sol)

pragma solidity ^0.8.0;

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}


// File contracts/PaymentLedger.sol


pragma solidity ^0.8.4;

contract PaymentLedger is Ownable{

    enum RequestStatus {PENDING, CANCELED, COMPLETED, FAILED}
    enum VaultStatus {INUSE, AVAILABLE}

    struct Request {
        string requester; // sender of the request.
        string token;
        uint amount; // amount of token to deposit/withdraw.
        string txid; // asset txid for the deposit/withdraw process.
        string vault;
        uint nonce; // serial number allocated for each request.
        uint timestamp; // time of the request creation.
        RequestStatus status; // status of the request.
    }

    struct VaultInfo {
        string vaultAddress;
        VaultStatus status;
    }

    Request[] public paymentRequests;
    Request[] public withdrawRequests;
    VaultInfo[] public vaults ;


    // mapping between a payment request hash and the corresponding request nonce.
    mapping(bytes32=>uint) public paymentRequestNonce;
    // txid => payment request hash
    mapping(string=>bytes32) public paymentRequestTxid;
    // vault address => vault nonce
    mapping(string=>uint) public vaultInfoNonce;

    mapping(bytes32=>uint) public withdrawRequestNonce;
    mapping(string=>bytes32) public withdrawRequestTxid;

    // vault address => (token ==> balance)
    mapping(string=>mapping(string=>uint)) public vaultTokenBalance;

    constructor() {}

    event PaymentRequestAdd(
        uint indexed nonce,
        string indexed requester,
        string indexed token,
        uint amount,
        string txid,
        string vault,
        uint timestamp,
        bytes32 requestHash
    );

    function AddPaymentRequest(
        string memory tokenAddress_,
        uint amount_,
        string memory txid_,
        string memory requester_
    )
    external
    onlyOwner
    returns (bool)
    {
        require(!isEmptyString(requester_), "invalid requester");

        uint nonce = paymentRequests.length;
        uint timestamp = getTimestamp();

        // get available vault
        uint vaultNonce = getAvailableVault();
        vaults[vaultNonce].status = VaultStatus.INUSE;

        Request memory request = Request({
        requester: requester_,
        token: tokenAddress_,
        amount: amount_,
        txid: txid_,
        vault: vaults[vaultNonce].vaultAddress,
        nonce: nonce,
        timestamp: timestamp,
        status: RequestStatus.PENDING
        });

        bytes32 requestHash = calcRequestHash(request);
        paymentRequestNonce[requestHash] = nonce;
        paymentRequestTxid[txid_] = requestHash;
        paymentRequests.push(request);

        emit PaymentRequestAdd(nonce, requester_, tokenAddress_, amount_, txid_,
            vaults[vaultNonce].vaultAddress, timestamp, requestHash);
        return true;
    }

    event PaymentRequestCancel(uint indexed nonce, string indexed txid, string indexed requester, bytes32 requestHash);

    function cancelPaymentRequest(string memory txid_, string memory requester_) external onlyOwner returns (bool) {
        require(!isEmptyString(txid_), "transaction id must not be empty");
        require(!isEmptyString(requester_), "requester must not be empty");
        uint nonce;
        Request memory request;
        bytes32 requestHash = paymentRequestTxid[txid_];
        (nonce, request) = getPendingPaymentRequest(requestHash);

        uint vaultNonce = vaultInfoNonce[request.vault];
        vaults[vaultNonce].status = VaultStatus.AVAILABLE;

        require(compareStrings(requester_, request.requester), "cancel sender is different than pending request initiator");
        paymentRequests[nonce].status = RequestStatus.CANCELED;

        emit PaymentRequestCancel(nonce, txid_, requester_, requestHash);
        return true;
    }

    event PaymentConfirmed(
        uint indexed nonce,
        string indexed requester,
        string indexed token,
        uint amount,
        string txid,
        string vault,
        uint timestamp,
        bytes32 requestHash
    );

    function confirmPaymentRequest(string memory txid_, string memory requester_) external onlyOwner returns (bool) {
        require(!isEmptyString(txid_), "transaction id must not be empty");
        require(!isEmptyString(requester_), "requester must not be empty");
        // change vault balance
        uint nonce;
        Request memory request;
        bytes32 requestHash = paymentRequestTxid[txid_];

        (nonce, request) = getPendingPaymentRequest(requestHash);


        uint vaultNonce = vaultInfoNonce[request.vault];
        vaults[vaultNonce].status = VaultStatus.AVAILABLE;
        vaultTokenBalance[request.vault][request.token] += request.amount;

        paymentRequests[nonce].status = RequestStatus.COMPLETED;

        emit PaymentConfirmed(
            request.nonce,
            request.requester,
            request.token,
            request.amount,
            request.txid,
            request.vault,
            request.timestamp,
            requestHash
        );
        return true;
    }

    event WithdrawRequestAdd(
        uint indexed nonce,
        string indexed requester,
        string indexed token,
        uint amount,
        string txid,
        string vault,
        uint timestamp,
        bytes32 requestHash
    );

    function AddWithdrawRequest(
        string memory tokenAddress_,
        uint amount_,
        string memory txid_,
        string memory vault_,
        string memory requester_
    )
    external
    onlyOwner
    returns (bool)
    {
        require(!isEmptyString(requester_), "invalid requester");
        require(!isEmptyString(txid_), "invalid transaction id");

        uint nonce = withdrawRequests.length;
        uint timestamp = getTimestamp();

        // get available vault
        uint vaultNonce = vaultInfoNonce[vault_];
        require(vaults[vaultNonce].status == VaultStatus.AVAILABLE, "the vault is not available");
        require(vaultTokenBalance[vault_][tokenAddress_] >= amount_,
            "The vault has insufficient amount to withdraw");

        vaults[vaultNonce].status = VaultStatus.INUSE;

        Request memory request = Request({
        requester: requester_,
        token: tokenAddress_,
        amount: amount_,
        txid: txid_,
        vault: vaults[vaultNonce].vaultAddress,
        nonce: nonce,
        timestamp: timestamp,
        status: RequestStatus.PENDING
        });

        bytes32 requestHash = calcRequestHash(request);
        withdrawRequestNonce[requestHash] = nonce;
        withdrawRequestTxid[txid_] = requestHash;
        withdrawRequests.push(request);

        emit WithdrawRequestAdd(nonce, requester_, tokenAddress_, amount_, txid_,
            vaults[vaultNonce].vaultAddress, timestamp, requestHash);
        return true;
    }

    event WithdrawRequestCancel(uint indexed nonce, string indexed txid, string indexed requester, bytes32 requestHash);

    function cancelWithdrawRequest(string memory txid_, string memory requester_) external onlyOwner returns (bool) {
        require(!isEmptyString(txid_), "transaction id must not be empty");
        require(!isEmptyString(requester_), "requester must not be empty");

        uint nonce;
        Request memory request;

        bytes32 requestHash = withdrawRequestTxid[txid_];
        (nonce, request) = getPendingWithdrawRequest(requestHash);

        uint vaultNonce = vaultInfoNonce[request.vault];
        vaults[vaultNonce].status = VaultStatus.AVAILABLE;

        require(compareStrings(requester_, request.requester), "cancel sender is different than pending request initiator");
        paymentRequests[nonce].status = RequestStatus.CANCELED;

        emit WithdrawRequestCancel(nonce, txid_, requester_, requestHash);
        return true;
    }

    event WithdrawConfirmed(
        uint indexed nonce,
        string indexed requester,
        string indexed token,
        uint amount,
        string txid,
        string vault,
        uint timestamp,
        bytes32 requestHash
    );

    function confirmWithdrawRequest(string memory txid_, string memory requester_) external onlyOwner returns (bool) {
        require(!isEmptyString(txid_), "transaction id must not be empty");
        require(!isEmptyString(requester_), "requester must not be empty");
        // change vault balance
        uint nonce;
        Request memory request;
        bytes32 requestHash = withdrawRequestTxid[txid_];

        (nonce, request) = getPendingWithdrawRequest(requestHash);


        uint vaultNonce = vaultInfoNonce[request.vault];
        vaults[vaultNonce].status = VaultStatus.AVAILABLE;
        vaultTokenBalance[request.vault][request.token] -= request.amount;

        withdrawRequests[nonce].status = RequestStatus.COMPLETED;

        emit WithdrawConfirmed(
            request.nonce,
            request.requester,
            request.token,
            request.amount,
            request.txid,
            request.vault,
            request.timestamp,
            requestHash
        );
        return true;
    }

    function resetPaymentVaults(string[] memory vaultAddresses_) external onlyOwner {
        while (vaults.length != 0) {
            vaults.pop();
        }

        for (uint i = 0; i < vaultAddresses_.length; i++) {
            require(!isEmptyString(vaultAddresses_[i]), "invalid vault address");
            vaults.push(VaultInfo({
                vaultAddress: vaultAddresses_[i],
                status: VaultStatus.AVAILABLE
            }));
        }
    }

    function getTimestamp() internal view returns (uint) {
        // timestamp is only used for data maintaining purpose, it is not relied on for critical logic.
        return block.timestamp; // solhint-disable-line not-rely-on-time
    }

    function compareStrings (string memory a, string memory b) internal pure returns (bool) {
        return (keccak256(abi.encodePacked(a)) == keccak256(abi.encodePacked(b)));
    }

    function isEmptyString (string memory a) internal pure returns (bool) {
        return (compareStrings(a, ""));
    }

    function calcRequestHash(Request memory request) internal pure returns (bytes32) {
        return keccak256(abi.encode(
                request.requester,
                request.amount,
                request.token,
                request.txid,
                request.vault,
                request.nonce,
                request.timestamp
            ));
    }

    function getAvailableVault() internal view returns (uint) {
        for (uint i = 0; i < vaults.length; i++) {
            VaultInfo storage vault = vaults[i];
            if (vault.status == VaultStatus.AVAILABLE) {
                return i;
            }
        }
        revert("No vault available");
    }

    function getPendingPaymentRequest(bytes32 requestHash) internal view returns (uint nonce, Request memory request) {
        require(requestHash != 0, "request hash is 0");
        nonce = paymentRequestNonce[requestHash];
        request = paymentRequests[nonce];
        validatePendingRequest(request, requestHash);
    }

    function validatePendingRequest(Request memory request, bytes32 requestHash) internal pure {
        require(request.status == RequestStatus.PENDING, "request is not pending");
        require(requestHash == calcRequestHash(request), "given request hash does not match a pending request");
    }

    function getPendingWithdrawRequest(bytes32 requestHash) internal view returns (uint nonce, Request memory request) {
        require(requestHash != 0, "request hash is 0");
        nonce = withdrawRequestNonce[requestHash];
        request = withdrawRequests[nonce];
        validatePendingRequest(request, requestHash);
    }

}