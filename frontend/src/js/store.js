
import { createStore } from 'redux';

import { initial } from './state';
import { RootReducer as root } from './root-reducer';

const store = createStore(root, initial());

export { store };
