import  express from "express";
import { createCar, getCar, testInvoke, testQuery } from "../controller/chaincodeController.js";
import { registerIdentity } from "../controller/identityController.js";

const router  = express.Router();

router.route('/api/register').post(registerIdentity)

router.route('/api/addCar').post(createCar)

router.route('/api/getCar').get(getCar)

router.route('/api/testInvoke').post(testInvoke)

router.route('/api/testQuery').post(testQuery)


export default router;
