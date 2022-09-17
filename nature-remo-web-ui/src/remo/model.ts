
/**
 * Specify \"power-off\" always if you want the air conditioner powered off. Empty means powered on.
 * @export
 * @enum {string}
 */
 export enum ACButton {
    Empty = <any> '',
    PowerOff = <any> 'power-off'
}

/**
 * 
 * @export
 * @interface AirCon
 */
export interface AirCon {
    /**
     * 
     * @type {any}
     * @memberof AirCon
     */
    range?: any;
    /**
     * 
     * @type {string}
     * @memberof AirCon
     */
    tempUnit?: AirCon.TempUnitEnum;
}

/**
 * @export
 * @namespace AirCon
 */
export namespace AirCon {
    /**
     * @export
     * @enum {string}
     */
    export enum TempUnitEnum {
        Empty = <any> '',
        C = <any> 'c',
        F = <any> 'f'
    }
}

/**
 * 
 * @export
 * @interface AirConParams
 */
export interface AirConParams {
    /**
     * 
     * @type {Temperature}
     * @memberof AirConParams
     */
    temp?: Temperature;
    /**
     * 
     * @type {OperationMode}
     * @memberof AirConParams
     */
    mode?: OperationMode;
    /**
     * 
     * @type {AirVolume}
     * @memberof AirConParams
     */
    vol?: AirVolume;
    /**
     * 
     * @type {AirDirection}
     * @memberof AirConParams
     */
    dir?: AirDirection;
    /**
     * 
     * @type {ACButton}
     * @memberof AirConParams
     */
    button?: ACButton;
}

/**
 * 
 * @export
 * @interface AirConRangeMode
 */
export interface AirConRangeMode {
    /**
     * 
     * @type {Array<Temperature>}
     * @memberof AirConRangeMode
     */
    temp?: Array<Temperature>;
    /**
     * 
     * @type {Array<AirVolume>}
     * @memberof AirConRangeMode
     */
    vol?: Array<AirVolume>;
    /**
     * 
     * @type {Array<AirDirection>}
     * @memberof AirConRangeMode
     */
    dir?: Array<AirDirection>;
}

/**
 * Empty means automatic.
 * @export
 * @enum {string}
 */
export enum AirDirection {
    Empty = <any> ''
}

/**
 * Empty means automatic. Numbers express the amount of volume. The range of AirVolumes which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
 * @export
 * @enum {string}
 */
export enum AirVolume {
    Empty = <any> '',
    Auto = <any> 'auto',
    _1 = <any> '1',
    _2 = <any> '2',
    _3 = <any> '3',
    _4 = <any> '4',
    _5 = <any> '5',
    _6 = <any> '6',
    _7 = <any> '7',
    _8 = <any> '8',
    _9 = <any> '9',
    _10 = <any> '10'
}

/**
 * 
 * @export
 * @interface Appliance
 */
export interface Appliance {
    /**
     * 
     * @type {Id}
     * @memberof Appliance
     */
    id?: Id;
    /**
     * 
     * @type {DeviceCore}
     * @memberof Appliance
     */
    device?: DeviceCore;
    /**
     * 
     * @type {ApplianceModel}
     * @memberof Appliance
     */
    model?: ApplianceModel;
    /**
     * 
     * @type {string}
     * @memberof Appliance
     */
    nickname?: string;
    /**
     * 
     * @type {Image}
     * @memberof Appliance
     */
    image?: Image;
    /**
     * 
     * @type {ApplianceType}
     * @memberof Appliance
     */
    type?: ApplianceType;
    /**
     * 
     * @type {AirConParams}
     * @memberof Appliance
     */
    settings?: AirConParams;
    /**
     * 
     * @type {AirCon}
     * @memberof Appliance
     */
    aircon?: AirCon;
    /**
     * 
     * @type {Array<Signal>}
     * @memberof Appliance
     */
    signals?: Array<Signal>;
    /**
     * 
     * @type {TV}
     * @memberof Appliance
     */
    tv?: TV;
    /**
     * 
     * @type {LIGHT}
     * @memberof Appliance
     */
    light?: LIGHT;
    /**
     * 
     * @type {SmartMeter}
     * @memberof Appliance
     */
    smartMeter?: SmartMeter;
}

/**
 * 
 * @export
 * @interface ApplianceModel
 */
export interface ApplianceModel {
    /**
     * 
     * @type {Id}
     * @memberof ApplianceModel
     */
    id?: Id;
    /**
     * 
     * @type {string}
     * @memberof ApplianceModel
     */
    manufacturer?: string;
    /**
     * 
     * @type {string}
     * @memberof ApplianceModel
     */
    remoteName?: string;
    /**
     * 
     * @type {string}
     * @memberof ApplianceModel
     */
    name?: string;
    /**
     * 
     * @type {Image}
     * @memberof ApplianceModel
     */
    image?: Image;
}

/**
 * 
 * @export
 * @interface ApplianceModelAndParam
 */
export interface ApplianceModelAndParam {
    /**
     * 
     * @type {ApplianceModel}
     * @memberof ApplianceModelAndParam
     */
    model?: ApplianceModel;
    /**
     * 
     * @type {AirConParams}
     * @memberof ApplianceModelAndParam
     */
    params?: AirConParams;
}

/**
 * Type of the appliance. \"AC\" (Air conditioner), \"TV\" and \"LIGHT\" are 1st class citizen appliance, which is included in our IRDB (InfraRed signals DataBase). The \"ApplianceModel\" stores meta data about the appliance. We provide AC specific UI. Everything else is \"IR\". We just learn the signals from the remote and store them, and when users tap the button on the smartphone app, our server sends it through Remo.
 * @export
 * @enum {string}
 */
export enum ApplianceType {
    AC = <any> 'AC',
    TV = <any> 'TV',
    LIGHT = <any> 'LIGHT',
    IR = <any> 'IR'
}

/**
 * 
 * @export
 * @interface Button
 */
export interface Button {
    /**
     * Name of button. It is used for \"POST /1/{applaince}/tv\" and \"POST /1/{appliance}/light\".
     * @type {string}
     * @memberof Button
     */
    name?: string;
    /**
     * 
     * @type {Image}
     * @memberof Button
     */
    image?: Image;
    /**
     * Label of button.
     * @type {string}
     * @memberof Button
     */
    label?: string;
}

/**
 * 
 * @export
 * @interface DateTime
 */
export interface DateTime {
}

/**
 * 
 * @export
 * @interface Device
 */
export interface Device {
    /**
     * 
     * @type {Id}
     * @memberof Device
     */
    id?: Id;
    /**
     * 
     * @type {string}
     * @memberof Device
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof Device
     */
    temperatureOffset?: number;
    /**
     * 
     * @type {number}
     * @memberof Device
     */
    humidityOffset?: number;
    /**
     * 
     * @type {Date}
     * @memberof Device
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof Device
     */
    updatedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof Device
     */
    firmwareVersion?: string;
    /**
     * 
     * @type {string}
     * @memberof Device
     */
    macAddress?: string;
    /**
     * 
     * @type {string}
     * @memberof Device
     */
    serialNumber?: string;
    /**
     * 
     * @type {any}
     * @memberof Device
     */
    newestEvents?: any;
}

/**
 * 
 * @export
 * @interface DeviceCore
 */
export interface DeviceCore {
    /**
     * 
     * @type {Id}
     * @memberof DeviceCore
     */
    id?: Id;
    /**
     * 
     * @type {string}
     * @memberof DeviceCore
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof DeviceCore
     */
    temperatureOffset?: number;
    /**
     * 
     * @type {number}
     * @memberof DeviceCore
     */
    humidityOffset?: number;
    /**
     * 
     * @type {Date}
     * @memberof DeviceCore
     */
    createdAt?: Date;
    /**
     * 
     * @type {Date}
     * @memberof DeviceCore
     */
    updatedAt?: Date;
    /**
     * 
     * @type {string}
     * @memberof DeviceCore
     */
    firmwareVersion?: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceCore
     */
    macAddress?: string;
    /**
     * 
     * @type {string}
     * @memberof DeviceCore
     */
    serialNumber?: string;
}

/**
 * 
 * @export
 * @interface Devices
 */
export interface Devices extends Array<Device> {
}

/**
 * The ECHONET lite properties fetched from the appliance. See \"Detailed Requirements for ECHONET Device Objects\" for more details. ref. https://echonet.jp/spec_object_rl_en/ 
 * @export
 * @interface EchonetLiteProperty
 */
export interface EchonetLiteProperty {
    /**
     * 
     * @type {string}
     * @memberof EchonetLiteProperty
     */
    name?: string;
    /**
     * ECHONET Property
     * @type {number}
     * @memberof EchonetLiteProperty
     */
    epc?: number;
    /**
     * 
     * @type {string}
     * @memberof EchonetLiteProperty
     */
    val?: string;
    /**
     * 
     * @type {Date}
     * @memberof EchonetLiteProperty
     */
    updatedAt?: Date;
}

/**
 * 
 * @export
 * @interface Id
 */
export interface Id {
}

/**
 * Basename of the image file included in the app. Ex: \"ico_ac_1\" 
 * @export
 * @interface Image
 */
export interface Image {
}

/**
 * 
 * @export
 * @interface LIGHT
 */
export interface LIGHT {
    /**
     * 
     * @type {LIGHTState}
     * @memberof LIGHT
     */
    state?: LIGHTState;
    /**
     * 
     * @type {Array<Button>}
     * @memberof LIGHT
     */
    buttons?: Array<Button>;
}

/**
 * 
 * @export
 * @interface LIGHTState
 */
export interface LIGHTState {
    /**
     * 
     * @type {string}
     * @memberof LIGHTState
     */
    brightness?: string;
    /**
     * 
     * @type {string}
     * @memberof LIGHTState
     */
    power?: LIGHTState.PowerEnum;
    /**
     * 
     * @type {string}
     * @memberof LIGHTState
     */
    lastButton?: string;
}

/**
 * @export
 * @namespace LIGHTState
 */
export namespace LIGHTState {
    /**
     * @export
     * @enum {string}
     */
    export enum PowerEnum {
        On = <any> 'on',
        Off = <any> 'off'
    }
}

/**
 * The range of OperationModes which the air conditioner accepts depends on the air conditioner model. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model.
 * @export
 * @enum {string}
 */
export enum OperationMode {
    Empty = <any> '',
    Cool = <any> 'cool',
    Warm = <any> 'warm',
    Dry = <any> 'dry',
    Blow = <any> 'blow',
    Auto = <any> 'auto'
}

/**
 * The reference key to SensorValue means \"te\" = temperature, \"hu\" = humidity, \"il\" = illumination, \"mo\" = movement. The val of \"mo\" is always 1 and when movement event is captured created_at is updated. 
 * @export
 * @interface SensorValue
 */
export interface SensorValue {
    /**
     * 
     * @type {number}
     * @memberof SensorValue
     */
    val?: number;
    /**
     * 
     * @type {Date}
     * @memberof SensorValue
     */
    createdAt?: Date;
}

/**
 * 
 * @export
 * @interface Signal
 */
export interface Signal {
    /**
     * 
     * @type {Id}
     * @memberof Signal
     */
    id?: Id;
    /**
     * 
     * @type {string}
     * @memberof Signal
     */
    name?: string;
    /**
     * 
     * @type {Image}
     * @memberof Signal
     */
    image?: Image;
}

/**
 * 
 * @export
 * @interface SmartMeter
 */
export interface SmartMeter {
    /**
     * 
     * @type {Array<EchonetLiteProperty>}
     * @memberof SmartMeter
     */
    echonetliteProperties?: Array<EchonetLiteProperty>;
}

/**
 * 
 * @export
 * @interface TV
 */
export interface TV {
    /**
     * 
     * @type {TVState}
     * @memberof TV
     */
    state?: TVState;
    /**
     * 
     * @type {Array<Button>}
     * @memberof TV
     */
    buttons?: Array<Button>;
}

/**
 * 
 * @export
 * @interface TVState
 */
export interface TVState {
    /**
     * 
     * @type {string}
     * @memberof TVState
     */
    input?: TVState.InputEnum;
}

/**
 * @export
 * @namespace TVState
 */
export namespace TVState {
    /**
     * @export
     * @enum {string}
     */
    export enum InputEnum {
        T = <any> 't',
        Bs = <any> 'bs',
        Cs = <any> 'cs'
    }
}

/**
 * The temperature in string format. The unit is described in Aircon object. The range of Temperatures which the air conditioner accepts depends on the air conditioner model and operation mode. Check the 'AirConRangeMode' information in the response for the range of the particular air conditioner model and operation mode.
 * @export
 * @interface Temperature
 */
export interface Temperature {
}

/**
 * 
 * @export
 * @interface User
 */
export interface User {
    /**
     * 
     * @type {Id}
     * @memberof User
     */
    id?: Id;
    /**
     * 
     * @type {string}
     * @memberof User
     */
    nickname?: string;
}

