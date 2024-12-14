import axios from 'axios';
import { alertService } from './alertor';

export const refreshService = {
    async refreshJWT(): Promise<boolean | null> {
        try {
            await axios.get("/api/auth/v1/refresh");
            return true;
        } catch (error: any) {
            if (error.response.data.need_login) {
                alertService.showAlert("You need to login bro 🎩", "info");
                return false;
            } else {
                console.error('Failed to refresh jwt', error);
                return null;
            }
        }
    }
}
