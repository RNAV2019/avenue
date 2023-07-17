import { initializeApp } from 'firebase/app';
import { firebaseConfig } from '../config/firebase/firebase-config';
import { getAuth } from 'firebase/auth';
// Initialize Firebase
const app = initializeApp(firebaseConfig);

// Initialize Firebase Authentication and get a reference to the service
export const auth = getAuth(app);
