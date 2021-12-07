// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import '@openzeppelin/contracts/token/ERC20/ERC20.sol';
import '@openzeppelin/contracts/access/Ownable.sol';

contract TokenB is Ownable, ERC20 {
  constructor() ERC20('TokenB', 'TokenB') {
  }

  function mint(uint amount) public {
    _mint(msg.sender, amount);
  }
}