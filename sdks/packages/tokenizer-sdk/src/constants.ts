import { TokenizerClientConfig } from './types';

const __TOKENIZER_URL__: string = process.env.TOKENIZER_URL;

export const CONFIG_DEFAULTS: TokenizerClientConfig = {
  host: __TOKENIZER_URL__,
  token: '',
  metaEncoding: 'base64',
  endpoints: {
    setMetadata: '/metadata/set',
    getMetadata: '/metadata/get',
    getToken: '/detokenize',
    setToken: '/tokenize',
  },
  headers: {
    auth: 'X-Lunasec-Session-Token'.toLowerCase(),
  },
};
