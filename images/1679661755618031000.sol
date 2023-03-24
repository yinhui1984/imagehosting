// Decompiled by library.dedaub.com
// 2023.03.08 11:37 UTC
// Compiled using the solidity compiler version 0.8.0


// Data structures and variables inferred from the use of storage instructions
uint256 _exchange; // STORAGE[0x4]
uint256 _one; // STORAGE[0x5]
uint256 stor_6; // STORAGE[0x6]
uint256 _three; // STORAGE[0x7]
uint256 stor_8; // STORAGE[0x8]
uint256 stor_9; // STORAGE[0x9]
mapping (uint256 => uint256) owner_a; // STORAGE[0xa]
mapping (uint256 => uint256) owner_b; // STORAGE[0xb]
mapping (uint256 => uint256) owner_c; // STORAGE[0xc]
mapping (uint256 => uint256) _operation; // STORAGE[0xd]
address _owner; // STORAGE[0x0] bytes 0 to 19
uint160 _lp; // STORAGE[0x1] bytes 0 to 19
address _usdt; // STORAGE[0x2] bytes 0 to 19
address stor_3_0_19; // STORAGE[0x3] bytes 0 to 19
address owner_e_0_19; // STORAGE[0xe] bytes 0 to 19
uint160 stor_f_0_19; // STORAGE[0xf] bytes 0 to 19
uint160 stor_10_0_19; // STORAGE[0x10] bytes 0 to 19


// Events
OwnershipTransferred(address, address);

function 0x1201() private { 
    require(stor_3_0_19.code.size);
    v0, v1 = stor_3_0_19.balanceOf(_lp).gas(msg.gas);
    require(v0); // checks call status, propagates error data on error
    MEM[64] = MEM[64] + (RETURNDATASIZE() + 31 & ~0x1f);
    require(MEM[64] + RETURNDATASIZE() - MEM[64] >= 32);
    0x2a44(v1);
    require(_usdt.code.size);
    v2, v3 = _usdt.balanceOf(_lp).gas(msg.gas);
    require(v2); // checks call status, propagates error data on error
    MEM[64] = MEM[64] + (RETURNDATASIZE() + 31 & ~0x1f);
    require(MEM[64] + RETURNDATASIZE() - MEM[64] >= 32);
    0x2a44(v3);
    v4 = 0x1af2(0xde0b6b3a7640000, v1);
    v5 = 0x1b6d(v3, v4);
    return v5;
}

function 0x1972(uint256 varg0, uint256 varg1, uint256 varg2) private { 
    v0 = address(varg1);
    v1 = address(varg2);
    if (this.balance >= 0) {
        if (v1.code.size > 0) {
            v2 = v3 = 0;
            while (v2 < 68) {
                MEM[MEM[64] + v2] = MEM[MEM[64] + 32 + v2];
                v2 = v2 + 32;
            }
            if (v2 > 68) {
                MEM[MEM[64] + 68] = 0;
            }
            v4, v5, v6, v7 = address(v1).transfer(v0, varg0).gas(msg.gas);
            if (RETURNDATASIZE() == 0) {
                v8 = v9 = 96;
            } else {
                v8 = v10 = new bytes[](RETURNDATASIZE());
                RETURNDATACOPY(v10.data, 0, RETURNDATASIZE());
            }
            if (!v4) {
                require(MEM[v8] <= 0v7, MEM[v8]);
                v11 = new array[](v12.length);
                v13 = v14 = 0;
                while (v13 < v12.length) {
                    v11[v13] = v12[v13];
                    v13 = v13 + 32;
                }
                if (v13 > v12.length) {
                    v11[32] = 0;
                }
                revert(Error(v11, v15, 'SafeERC20: low-level call failed'));
            } else {
                if (MEM[v8] > 0) {
                    require(32 + v8 + MEM[v8] - (32 + v8) >= 32);
                    0x2a2d(MEM[32 + v8 + 0]);
                    require(MEM[32 + v8 + 0], Error('SafeERC20: ERC20 operation did not succeed'));
                }
                return ;
            }
        } else {
            MEM[4 + MEM[64] + 0] = 32;
            revert(Error('Address: call to non-contract'));
        }
    } else {
        MEM[4 + MEM[64] + 0] = 32;
        revert(Error('Address: insufficient balance for call'));
    }
}

function _SafeAdd(uint256 varg0, uint256 varg1) private { 
    require(varg1 <= 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff - varg0, Panic(17));
    v0 = varg1 + varg0;
    require(v0 >= varg1, Error('SafeMath: addition overflow'));
    return v0;
}

function 0x1a69(uint256 varg0, uint256 varg1, uint256 varg2, uint256 varg3) private { 
    v0 = address(varg2);
    v1 = address(varg1);
    v2 = address(varg3);
    if (this.balance >= 0) {
        if (v2.code.size > 0) {
            v3 = v4 = 0;
            while (v3 < 100) {
                MEM[MEM[64] + v3] = MEM[MEM[64] + 32 + v3];
                v3 = v3 + 32;
            }
            if (v3 > 100) {
                MEM[MEM[64] + 100] = 0;
            }
            v5, v6, v7, v8 = address(v2).transferFrom(v0, v1, varg0).gas(msg.gas);
            if (RETURNDATASIZE() == 0) {
                v9 = v10 = 96;
            } else {
                v9 = v11 = new bytes[](RETURNDATASIZE());
                RETURNDATACOPY(v11.data, 0, RETURNDATASIZE());
            }
            if (!v5) {
                require(MEM[v9] <= 0v8, MEM[v9]);
                v12 = new array[](v13.length);
                v14 = v15 = 0;
                while (v14 < v13.length) {
                    v12[v14] = v13[v14];
                    v14 = v14 + 32;
                }
                if (v14 > v13.length) {
                    v12[32] = 0;
                }
                revert(Error(v12, v16, 'SafeERC20: low-level call failed'));
            } else {
                if (MEM[v9] > 0) {
                    require(32 + v9 + MEM[v9] - (32 + v9) >= 32);
                    0x2a2d(MEM[32 + v9 + 0]);
                    require(MEM[32 + v9 + 0], Error('SafeERC20: ERC20 operation did not succeed'));
                }
                return ;
            }
        } else {
            MEM[4 + MEM[64] + 0] = 32;
            revert(Error('Address: call to non-contract'));
        }
    } else {
        MEM[4 + MEM[64] + 0] = 32;
        revert(Error('Address: insufficient balance for call'));
    }
}

function 0x1af2(uint256 varg0, uint256 varg1) private { 
    if (varg1 != 0) {
        require(!(varg1 & (varg0 > 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff / varg1)), Panic(17));
        v0 = v1 = varg1 * varg0;
        v2 = _SafeDiv(v1, varg1);
        require(v2 == varg0, Error('SafeMath: multiplication overflow'));
    } else {
        v0 = v3 = 0;
    }
    return v0;
}

function 0x1b6d(uint256 varg0, uint256 varg1) private { 
    require(varg0 > 0, Error('SafeMath: division by zero'));
    v0 = _SafeDiv(varg1, varg0);
    return v0;
}

function _SafeSub(uint256 varg0, uint256 varg1) private { 
    require(varg0 <= varg1, Error('SafeMath: subtraction overflow'));
    require(varg1 >= varg0, Panic(17));
    return varg1 - varg0;
}

function getUsdt(address varg0, uint256 varg1) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 64);
    0x2a16(varg0);
    0x2a44(varg1);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    0x1972(varg1, varg0, _usdt);
    return 1;
}

function operation(address varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    return 0xff & _operation[varg0];
}

function 0x20c995a2(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    stor_10_0_19 = varg0;
}

function usdt() public payable { 
    return _usdt;
}

function lp() public payable { 
    return _lp;
}

function getToken(address varg0, uint256 varg1) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 64);
    0x2a16(varg0);
    0x2a44(varg1);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    0x1972(varg1, varg0, stor_3_0_19);
    return 1;
}

function _SafeDiv(uint256 varg0, uint256 varg1) private { 
    require(varg1, Panic(18));
    return varg0 / varg1;
}

function 0x2a16(uint256 varg0) private { 
    require(varg0 == address(varg0));
    return ;
}

function 0x2a2d(uint256 varg0) private { 
    require(varg0 == varg0);
    return ;
}

function 0x2a44(uint256 varg0) private { 
    require(varg0 == varg0);
    return ;
}

function () public payable { 
    revert();
}

function three() public payable { 
    return _three;
}

function exchange(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a44(varg0);
    require(msg.sender.code.size <= 0, Error('no isContract'));
    require(varg0 >= stor_9, Error('num >= propor'));
    v0 = _SafeAdd(varg0, owner_a[msg.sender]);
    owner_a[msg.sender] = v0;
    0x1a69(varg0, stor_10_0_19, msg.sender, _usdt);
    require(owner_e_0_19.code.size);
    v1, v2, v3, v4 = owner_e_0_19.staticcall(0xbd52993b, msg.sender).gas(msg.gas);
    require(v1); // checks call status, propagates error data on error
    MEM[64] = MEM[64] + (RETURNDATASIZE() + 31 & ~0x1f);
    require(MEM[64] + RETURNDATASIZE() - MEM[64] >= 96);
    0x2a16(v2);
    0x2a16(v3);
    0x2a16(v4);
    v5 = 0x1201();
    v6 = 0x1af2(varg0, v5);
    v7 = 0x1b6d(0xde0b6b3a7640000, v6);
    v8 = v9 = 0;
    if (address(v2) != 0) {
        v10 = 0x1af2(_one, v7);
        v11 = 0x1b6d(_exchange, v10);
        v8 = v12 = _SafeAdd(v11, v9);
        v13 = _SafeAdd(v11, owner_b[address(v2)]);
        owner_b[address(v2)] = v13;
    }
    if (address(v3) != 0) {
        v14 = 0x1af2(stor_6, v7);
        v15 = 0x1b6d(_exchange, v14);
        v8 = v16 = _SafeAdd(v15, v8);
        v17 = _SafeAdd(v15, owner_b[address(v3)]);
        owner_b[address(v3)] = v17;
    }
    if (address(v4) != 0) {
        v18 = 0x1af2(_three, v7);
        v19 = 0x1b6d(_exchange, v18);
        v8 = v20 = _SafeAdd(v19, v8);
        v21 = _SafeAdd(v19, owner_b[address(v4)]);
        owner_b[address(v4)] = v21;
    }
    v22 = 0x1af2(stor_8, v7);
    v23 = 0x1b6d(_exchange, v22);
    v24 = _SafeAdd(v23, v8);
    0x1972(v23, stor_f_0_19, stor_3_0_19);
    v25 = _SafeSub(v24, v7);
    0x1972(v25, msg.sender, stor_3_0_19);
}

function 0x5fdf05d7() public payable { 
    return stor_6;
}

function renounceOwnership() public payable { 
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    emit OwnershipTransferred(_owner, 0);
    _owner = 0;
}

function 0x812a3acf() public payable { 
    v0 = _SafeSub(owner_c[msg.sender], owner_b[msg.sender]);
    if (v0 > 0) {
        v1 = _SafeAdd(v0, owner_c[msg.sender]);
        owner_c[msg.sender] = v1;
        0x1972(v0, msg.sender, stor_3_0_19);
    }
}

function 0x85ce2405() public payable { 
    return stor_9;
}

function owner() public payable { 
    return _owner;
}

function one() public payable { 
    return _one;
}

function 0x961bdfbf() public payable { 
    v0 = 0x1201();
    return v0;
}

function getTokens(address varg0, address varg1, uint256 varg2) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 96);
    0x2a16(varg0);
    0x2a16(varg1);
    0x2a44(varg2);
    require(0xff & _operation[msg.sender], Error('no operation'));
    0x1a69(varg2, varg1, varg0, _usdt);
}

function 0xaddb3028(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a44(varg0);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    stor_9 = varg0;
}

function 0xb2d34d55(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    return owner_c[varg0];
}

function 0xbb11049f(uint256 varg0, uint256 varg1) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 64);
    0x2a16(varg0);
    0x2a2d(varg1);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    _operation[address(varg0)] = varg1 | ~0xff & _operation[address(varg0)];
}

function 0xbee3157e() public payable { 
    return stor_10_0_19;
}

function 0xcf1c9453() public payable { 
    return stor_f_0_19;
}

function 0xd42568f7(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    stor_f_0_19 = varg0;
}

function 0xe176895e(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    return owner_b[varg0];
}

function 0xe274a7bc() public payable { 
    return owner_e_0_19;
}

function 0xe6db2713() public payable { 
    return stor_3_0_19;
}

function transferOwnership(address varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    require(_owner == msg.sender, Error('Ownable: caller is not the owner'));
    require(varg0 != 0, Error('Ownable: new owner is the zero address'));
    emit OwnershipTransferred(_owner, varg0);
    _owner = varg0;
}

function 0xfd19016c() public payable { 
    return stor_8;
}

function 0xfdc65aa7(uint256 varg0) public payable { 
    require(4 + (msg.data.length - 4) - 4 >= 32);
    0x2a16(varg0);
    return owner_a[varg0];
}

// Note: The function selector is not present in the original solidity code.
// However, we display it for the sake of completeness.

function __function_selector__(bytes4 function_selector) public payable { 
    MEM[64] = 128;
    require(!msg.value);
    if (msg.data.length < 4) {
        ();
    } else {
        v0 = function_selector >> 224;
        if (0x961bdfbf > v0) {
            if (0x53556559 > v0) {
                if (0x2f48ab7d > v0) {
                    if (0x6bdd7d2 == v0) {
                        getUsdt(address,uint256);
                    } else if (0x12c57c32 == v0) {
                        operation(address);
                    } else {
                        require(0x20c995a2 == v0);
                        0x20c995a2();
                    }
                } else if (0x2f48ab7d == v0) {
                    usdt();
                } else if (0x313c06a0 == v0) {
                    lp();
                } else if (0x43d7cce6 == v0) {
                    getToken(address,uint256);
                } else {
                    require(0x45caa117 == v0);
                    three();
                }
            } else if (0x812a3acf > v0) {
                if (0x53556559 == v0) {
                    exchange(uint256);
                } else if (0x5fdf05d7 == v0) {
                    0x5fdf05d7();
                } else {
                    require(0x715018a6 == v0);
                    renounceOwnership();
                }
            } else if (0x812a3acf == v0) {
                0x812a3acf();
            } else if (0x85ce2405 == v0) {
                0x85ce2405();
            } else if (0x8da5cb5b == v0) {
                owner();
            } else {
                require(0x901717d1 == v0);
                one();
            }
        } else if (0xd42568f7 > v0) {
            if (0xb2d34d55 > v0) {
                if (0x961bdfbf == v0) {
                    0x961bdfbf();
                } else if (0xab7e8b5d == v0) {
                    getTokens(address,address,uint256);
                } else {
                    require(0xaddb3028 == v0);
                    0xaddb3028();
                }
            } else if (0xb2d34d55 == v0) {
                0xb2d34d55();
            } else if (0xbb11049f == v0) {
                0xbb11049f();
            } else if (0xbee3157e == v0) {
                0xbee3157e();
            } else {
                require(0xcf1c9453 == v0);
                0xcf1c9453();
            }
        } else if (0xe6db2713 > v0) {
            if (0xd42568f7 == v0) {
                0xd42568f7();
            } else if (0xe176895e == v0) {
                0xe176895e();
            } else {
                require(0xe274a7bc == v0);
                0xe274a7bc();
            }
        } else if (0xe6db2713 == v0) {
            0xe6db2713();
        } else if (0xf2fde38b == v0) {
            transferOwnership(address);
        } else if (0xfd19016c == v0) {
            0xfd19016c();
        } else {
            require(0xfdc65aa7 == v0);
            0xfdc65aa7();
        }
    }
}
