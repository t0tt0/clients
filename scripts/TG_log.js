var logger = {
    prefix: (id) => {
        return (new Date()) + " [Worker_" + id + "]";
    },
    consoleLog: (id, message) => {
        console.log(logger.prefix(id) + " " + message);
    },
    stringLog: (id, message) => {
        return logger.prefix(id) + " " + message;
    },
};

module.exports = logger;
