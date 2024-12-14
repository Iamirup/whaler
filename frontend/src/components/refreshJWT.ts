import axios from 'axios';
import { useRouter } from 'vue-router';
import { alertService } from './alertor';

const router = useRouter();

export const refreshJWT = async (): Promise<boolean> => {
    await axios.get("/api/auth/v1/refresh")
    .then(response => {})
    .catch(error => {
        console.error(error);
        if (error.response.data.need_login){
            alertService.showAlert("You need to login bro 🎩", "info");
            return false;
        } else {
            console.error('Failed to refresh jwt', error);
        }
    });

    return true;
}