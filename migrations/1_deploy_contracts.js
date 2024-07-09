const price = artifacts.require("Price");

module.exports = function(deployer) {
	deployer.deploy(price);
};
