import axios, { AxiosResponse } from 'axios';
import { BackendConfig } from '../config';
import { Appliance } from './model';

const config = BackendConfig;
const baseURL = config.scheme == undefined ? 
    window.location.href + '/api/v1' :
    `${config.scheme}://${config.host}:${config.port}/api/v1`

console.log(`API base URL: ${baseURL}`);

const instance = axios.create({
    baseURL: baseURL,
    timeout: 30 * 1000
});

export async function getAppliances(force: boolean): Promise<Appliance[]> {
        const response = await instance.get(`/appliances?force=${force}`);
        const appliances: Appliance[] = await response['data'];
        return appliances;
}

export async function sendSignal(signalID: string) {
    await instance.post(`/signals/${signalID}/send`) 
}