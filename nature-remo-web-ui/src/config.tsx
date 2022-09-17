const Config = () => {
    return {
        backend: {
            scheme: process.env.REACT_APP_BACKEND_SCHEME,
            host: process.env.REACT_APP_BACKEND_HOST,
            port: process.env.REACT_APP_BACKEND_PORT,
            suffix: process.env.REACT_APP_BACKEND_SUFFIX,
        }
    }
}

console.log(`env: ${process.env}`)
console.log(`backend server: ${process.env.BACKEND_SCHEME}://${process.env.BACKEND_HOST}:${process.env.BACKEND_PORT}`);

export const BackendConfig = Config().backend