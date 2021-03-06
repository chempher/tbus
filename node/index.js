// for data types
require('./gen/tbus/bus_pb.js');
require('./gen/tbus/led_pb.js');
require('./gen/tbus/motor_pb.js');
require('./gen/tbus/servo_pb.js');

module.exports = {
    Device:     require('./lib/device.js'),
    Controller: require('./lib/control.js'),
    Bus:    require('./lib/bus.js'),
    Master: require('./lib/master.js'),
    RemoteBusPort:    require('./lib/remotebusport.js'),
    RemoteDeviceHost: require('./lib/remotedevhost.js'),
    SocketConnector: require('./lib/socketconn.js'),

    protocol: require('./lib/protocol.js'),

    BusDev: require('./gen/tbus/bus_tbusdev.js').BusDev,
    BusCtl: require('./gen/tbus/bus_tbusdev.js').BusCtl,
    LEDDev: require('./gen/tbus/led_tbusdev.js').LEDDev,
    LEDCtl: require('./gen/tbus/led_tbusdev.js').LEDCtl,
    MotorDev: require('./gen/tbus/motor_tbusdev.js').MotorDev,
    MotorCtl: require('./gen/tbus/motor_tbusdev.js').MotorCtl,
    ServoDev: require('./gen/tbus/servo_tbusdev.js').ServoDev,
    ServoCtl: require('./gen/tbus/servo_tbusdev.js').ServoCtl,
};

// for extensions
require('./lib/deviceinfo.js');
require('./lib/ctls/led.js');
require('./lib/ctls/motor.js');
require('./lib/ctls/servo.js');
