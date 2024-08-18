// script/DeployMockEAS.s.sol
// SPDX-License-Identifier: MIT
pragma solidity 0.8.24;

import "forge-std/src/Script.sol";
import "./MockEAS.sol";

contract DeployMockEAS is Script {
    function run() external {
        vm.startBroadcast();
        MockEAS mockEAS = new MockEAS();
        vm.stopBroadcast();
        console.log("MockEAS deployed at", address(mockEAS));
    }
}
