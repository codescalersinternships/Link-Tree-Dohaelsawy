/// <reference types="vite/client" />

declare namespace NodeJS {
    interface ImportMetaEnv {
        VITE_AXIOS_API: string;
        VITE_APP_API: string;
    }
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}