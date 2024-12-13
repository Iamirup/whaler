<template>
	
<div class="assets">
      <div class="container" :class="{ 'right-panel-active': isRightPanelActive }">
        
        <div class="container__form container--signup">
            <form @submit.prevent="register" action="#" class="form" id="form1">
                <h2 class="form__title">Register</h2>
				<div class="input-group">
					<input type="text" required="" name="text" v-model="registerData.username"  autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Username</label>
				</div>
				<div class="input-group my-1">
					<input type="text" required="" name="text" v-model="registerData.email"  autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Email</label>
				</div>
				<div class="input-group my-1">
					<input type="password" required="" name="text" v-model="registerData.password"  autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Password</label>
				</div>
				<div class="input-group my-1">
					<input type="password" required="" name="text" v-model="registerData.confirm_password"  autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Confirm Password</label>
				</div>
                <button class="btn w-12 flex justify-center">Register</button>
            </form>
        </div>
    
        
        <div class="container__form container--signin">
            <form @submit.prevent="login" action="#" class="form" id="form2">
                <h2 class="form__title">Login</h2>
				<div class="input-group my-1">
					<input type="text" required="" name="text" v-model="loginData.identifier"  autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Email or Username</label>
				</div>
				<div class="input-group">
					<input type="password" required="" name="text" v-model="loginData.password" autocomplete="off" class="input rounded-xl h-11">
					<label class="user-label">Password</label>
				</div>
                <a href="#" class="link">Forgot your password?</a>
                <button class="btn w-12 flex justify-center">Login</button>
            </form>
        </div>
    
        
        <div class="container__overlay">
            <div class="overlay">
                <div class="overlay__panel overlay--left">
                    <button class="btn" id="signIn" @click="togglePanel(false)">Login</button>
                </div>
                <div class="overlay__panel overlay--right">
                    <button class="btn" id="signUp" @click="togglePanel(true)">Register</button>
                </div>
            </div>
        </div>
    </div>
	
</div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';
import { alertService } from '../alertor';
import { useRouter } from 'vue-router';
const router = useRouter();

const axioser = axios.create({
	baseURL: 'https://whaler.ir'
});

interface RegisterData {
	username: string;
	email: string;
	password: string;
	confirm_password: string; 
}

interface LoginData {
	identifier: string;
	password: string;
}

export default defineComponent({
  name: 'AuthComponent',
  data() {
	return {
		registerData: {
			username: '',
			email: '',
			password: '',
			confirm_password: '',
		} as RegisterData,
		loginData: {
			identifier: '',
			password: '',
		} as LoginData,
	};
  },
  methods: {
	async login() {
		const isEmail = this.validateEmail(this.loginData.identifier);
		const loginPayload = isEmail 
			? { email: this.loginData.identifier, password: this.loginData.password }
			: { username: this.loginData.identifier, password: this.loginData.password }


		await axioser.post("/api/auth/v1/login", loginPayload)
		.then(response => {
			console.log(response)
			this.setCookie("access_token", response.data.access_token);
			alertService.showAlert("Successful login", "success");
			router.push('/eventor');
		})
		.catch(error => {
			let alertErrorMessage = ""
			for (const obj of error.response.data.errors) {
				alertErrorMessage += "- " + obj.message + ".\n";
			}
			alertErrorMessage = alertErrorMessage.trim();
			alertService.showAlert(alertErrorMessage, "error");
		});

		this.loginData = {
			identifier: '',
			password: '',
		}
	},
	async register() {
		await axioser.post("/api/auth/v1/register", this.registerData)
		.then(response => {
			this.setCookie("access_token", response.data.access_token);
			alertService.showAlert("Successful regestration", "success");
			router.push('/eventor');
		})
		.catch(error => {
			let alertErrorMessage = ""
			for (const obj of error.response.data.errors) {
				alertErrorMessage += "- " + obj.message + ".\n";
			}
			alertErrorMessage = alertErrorMessage.trim();
			alertService.showAlert(alertErrorMessage, "error");
		});

		this.registerData = {
			username: '',
			email: '',
			password: '',
			confirm_password: '',
		}
	},
	setCookie(name: string, value: string) {
		document.cookie = `${name}=${value}; path=/; HttpOnly;`
	},
	validateEmail(identifier: string): boolean {
		const re = /\S+@\S+\.\S+/;
		return re.test(identifier);
	}
  },
  setup() {
    const isRightPanelActive = ref(false);

    const togglePanel = (isActive: boolean) => {
      isRightPanelActive.value = isActive;
    };

    return {
      isRightPanelActive,
      togglePanel
    };
  }
})
</script>

<style>
 
.input-group {
 position: relative;
 width: 275px;
}

.input {
 border: solid 1.5px #9e9e9e;
 border-radius: 1rem;
 background: none;
 padding: 1rem;
 font-size: 1rem;
 color: #131212;
 transition: border 150ms cubic-bezier(0.4,0,0.2,1);
}

.user-label {
 position: absolute;
 left: 15px;
 color: #adabab;
 pointer-events: none;
 transform: translateY(1rem);
 transition: 150ms cubic-bezier(0.4,0,0.2,1);
}

.input:focus, input:valid {
 outline: none;
 border: 1.5px solid #1a73e8;
}

.input:focus ~ label, input:valid ~ label {
 transform: translateY(-50%) scale(0.8);
 /* background-color: #212121; */
 padding: 0 .2em;
 color: #9e9e9e;
}





:root {
	/* COLORS */
	--white: #e9e9e9;
	--gray: #333;
	--blue: #0367a6;
	--lightblue: #008997;

	/* RADII */
	--button-radius: 0.7rem;

	/* SIZES */
	--max-width: 758px;
	--max-height: 420px;

	font-size: 16px;
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
		Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
}

.assets {
	align-items: center;
	background-color: var(--white);
	background: url("../../assets/dth.jpg");
	background-attachment: fixed;
	background-position: center;
	background-repeat: no-repeat;
	background-size: cover;
	display: grid;
	height: 100vh;
	place-items: center;
    overflow: hidden;
}

.form__title {
	font-weight: 300;
	margin: 0;
	margin-bottom: 1.25rem;
}

.link {
	color: var(--gray);
	font-size: 0.9rem;
	margin: 1.5rem 0;
	text-decoration: none;
}

.container {
	background-color: var(--white);
	border-radius: var(--button-radius);
	box-shadow: 0 0.9rem 1.7rem rgba(0, 0, 0, 0.25),
		0 0.7rem 0.7rem rgba(0, 0, 0, 0.22);
	height: var(--max-height);
	max-width: var(--max-width);
	overflow: hidden;
	position: relative;
	width: 100%;
}

.container__form {
	height: 100%;
	position: absolute;
	top: 0;
	transition: all 0.6s ease-in-out;
}

.container--signin {
	left: 0;
	width: 50%;
	z-index: 2;
}

.container.right-panel-active .container--signin {
	transform: translateX(100%);
}

.container--signup {
	left: 0;
	opacity: 0;
	width: 50%;
	z-index: 1;
}

.container.right-panel-active .container--signup {
	animation: show 0.6s;
	opacity: 1;
	transform: translateX(100%);
	z-index: 5;
}

.container__overlay {
	height: 100%;
	left: 50%;
	overflow: hidden;
	position: absolute;
	top: 0;
	transition: transform 0.6s ease-in-out;
	width: 50%;
	z-index: 100;
}

.container.right-panel-active .container__overlay {
	transform: translateX(-100%);
}

.overlay {
	background-color: var(--lightblue);
	background: url("../../assets/dth.jpg");
	background-attachment: fixed;
	background-position: center;
	background-repeat: no-repeat;
	background-size: cover;
	height: 100%;
	left: -100%;
	position: relative;
	transform: translateX(0);
	transition: transform 0.6s ease-in-out;
	width: 200%;
}

.container.right-panel-active .overlay {
	transform: translateX(50%);
}

.overlay__panel {
	align-items: center;
	display: flex;
	flex-direction: column;
	height: 100%;
	justify-content: center;
	position: absolute;
	text-align: center;
	top: 0;
	transform: translateX(0);
	transition: transform 0.6s ease-in-out;
	width: 50%;
}

.overlay--left {
	transform: translateX(-20%);
}

.container.right-panel-active .overlay--left {
	transform: translateX(0);
}

.overlay--right {
	right: 0;
	transform: translateX(0);
}

.container.right-panel-active .overlay--right {
	transform: translateX(20%);
}

.btn {
	background-color: var(--blue);
	background-image: linear-gradient(90deg, var(--blue) 0%, var(--lightblue) 74%);
	border-radius: 20px;
	border: 1px solid var(--blue);
	color: var(--white);
	cursor: pointer;
	font-size: 0.8rem;
	font-weight: bold;
	letter-spacing: 0.1rem;
	padding: 0.9rem 4rem;
	text-transform: uppercase;
	transition: transform 80ms ease-in;
}

.form > .btn {
	margin-top: 1.5rem;
}

.btn:active {
	transform: scale(0.95);
}

.btn:focus {
	outline: none;
}

.form {
	background-color: var(--white);
	display: flex;
	align-items: center;
	justify-content: center;
	flex-direction: column;
	padding: 0 3rem;
	height: 100%;
	text-align: center;
}

.input {
	background-color: #fff;
	border: none;
	padding: 0.9rem 0.9rem;
	margin: 0.5rem 0;
	width: 100%;
}
</style>