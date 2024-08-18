// scripts/deploy.js
async function main() {
    const MockEAS = await ethers.getContractFactory("MockEAS");
    const mockEAS = await MockEAS.deploy();
    await mockEAS.deployed();
    console.log("MockEAS deployed to:", mockEAS.address);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });
