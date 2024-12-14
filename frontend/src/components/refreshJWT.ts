import axios from 'axios';
import { alertService } from './alertor';

export const refreshJWT = async (): Promise<boolean> => {
    await axios.get("/api/auth/v1/refresh")
    .then(response => {})
    .catch(error => {
        if (error.response.data.need_login){
            alertService.showAlert("You need to login bro 🎩", "info");
            return false;
        } else {
            console.error('Failed to refresh jwt', error);
        }
    });

    return true;
}