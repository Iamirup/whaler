import axios from 'axios';

export const refreshJWT = async () => {
    await axios.get("/api/auth/v1/refresh")
    .then(response => {})
    .catch(error => {
        console.error('Failed to refresh jwt', error);
    });
}