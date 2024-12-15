import axios from 'axios';

export const adminService = {
    async isAdmin(): Promise<boolean | null> {
        try {
            const response = await axios.get("/api/auth/v1/is_admin");
            return response.data.is_admin
        } catch (error: any) {
            console.error('Failed to use api', error);
            return null;
        }
    }
}
