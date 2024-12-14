import axios from 'axios';
import { useRouter } from 'vue-router';
import { alertService } from './alertor';

const router = useRouter();

export const refreshJWT = async () => {
    await axios.get("/api/auth/v1/refresh")
    .then(response => {})
    .catch(error => {
        if (error.response.data.need_login){
            alertService.showAlert("You need to login bro 🎩", "info");
            router.push('/login');
            return false;
        } else {
            console.error('Failed to refresh jwt', error);
        }
    });

    return true;
}