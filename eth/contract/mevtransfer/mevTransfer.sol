// SPDX-License-Identifier: undefined
pragma solidity 0.6.6;
interface IChiToken {
    function freeFromUpTo(address from, uint256 value) external;
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);
}

contract MevTransfer{
    IChiToken constant public chi = IChiToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);

    modifier discountCHI {
        uint gasStart = gasleft();
        _;
        uint gasSpent = 21000 + gasStart - gasleft() + 16 * msg.data.length;
        uint chiToUse = (gasSpent + 14154) / 41947;
        if (chiToUse > 0 && chi.balanceOf(msg.sender) >= chiToUse && chi.allowance(msg.sender,address(this)) >= chiToUse){
            chi.freeFromUpTo(msg.sender, (gasSpent + 14154) / 41947);
        }
    }

    function transfer(uint bribe)  external payable discountCHI {
        block.coinbase.transfer(bribe);
        msg.sender.transfer(address(this).balance);
    }
}