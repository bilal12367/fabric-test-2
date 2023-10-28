import { DefaultEventHandlerStrategies, Gateway, Wallets } from 'fabric-network';
import { fileURLToPath } from 'url';
import { dirname } from 'path';
import { getCCPJsonFile, getWalletPath } from '../utils/helper.js';
import logger from '../logger/logger.js';


export const invoke = async (orgname, username, funcname, args) => {

    const ccp = await getCCPJsonFile(orgname) //JSON.parse(ccpJSON);

    // Create a new file system based wallet for managing identities.
    const walletPath = await getWalletPath(orgname) //path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    // Check to see if we've already enrolled the user.
    let identity = await wallet.get(username);
    if (!identity) {
        logger.fatal("Identity not registered !!! ", username)
        return
    }



    const connectOptions = {
        wallet, identity: username, discovery: { enabled: true, asLocalhost: false },
        eventHandlerOptions: {
            commitTimeout: 100,
            strategy: DefaultEventHandlerStrategies.NETWORK_SCOPE_ALLFORTX
        }
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, connectOptions);

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork('mychannel');
    logger.debug("args: ", args)
    const contract = network.getContract('fabcar');
    const result = await contract.submitTransaction(funcname, ...args);
    return result;
}

export const query = async (orgname, username, funcname, args) => {
    const ccp = await getCCPJsonFile(orgname) //JSON.parse(ccpJSON);

    // Create a new file system based wallet for managing identities.
    const walletPath = await getWalletPath(orgname) //path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    // Check to see if we've already enrolled the user.
    let identity = await wallet.get(username);
    if (!identity) {
        logger.fatal("Identity not registered !!! ", username)
        return
    }



    const connectOptions = {
        wallet, identity: username, discovery: { enabled: true, asLocalhost: false },
        eventHandlerOptions: {
            commitTimeout: 100,
            strategy: DefaultEventHandlerStrategies.NETWORK_SCOPE_ALLFORTX
        }
    }

    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, connectOptions);

    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork('mychannel');

    const contract = network.getContract('fabcar');

    // const result = await contract.submitTransaction('createCar', args[0], args[1], args[2], args[3], args[4]);
    const result = await contract.evaluateTransaction(funcname, ...args)
    logger.info("Successfully Query!!!")
    logger.debug("Result From Query: ", result)
    const resultJSON = JSON.parse(result.toString())
    return resultJSON;
}