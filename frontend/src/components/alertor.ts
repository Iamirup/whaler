import Swal, { type SweetAlertIcon } from 'sweetalert2';

export const alertService = {
    showAlert(text: string, icon: SweetAlertIcon) {
        const Toast = Swal.mixin({
            toast: true,
            position: "top-end",
            showConfirmButton: false,
            timer: 5000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.onmouseenter = Swal.stopTimer;
                toast.onmouseleave = Swal.resumeTimer;
            }
        });
        Toast.fire({
            icon: icon,
            title: text
        });
    }
};