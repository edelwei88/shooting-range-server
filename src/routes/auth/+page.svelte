<script lang="ts">
	import { invoke } from '@tauri-apps/api/tauri';
	import { Button, Input, Label } from 'flowbite-svelte';
	invoke('href_auth');

	let valueLogin: String;
	let valuePassword: String;

	async function auth() {
		if ((await invoke('auth', { login: valueLogin, password: valuePassword })) == true) {
			document.location.href = '/main';
		}
	}

	function enter_press(event: KeyboardEvent) {
		if (event.key == 'Enter') {
			auth();
		}
	}
</script>

<div class="container max-w-md px-5 mt-6">
	<div class="mb-3">
		<Label for="login-input" class="block mb-2">Логин</Label>
		<Input id="login-input" size="lg" placeholder="Введите ваш логин" bind:value={valueLogin} />
	</div>

	<div class="mb-6">
		<Label for="password-input" class="block mb-2">Пароль</Label>
		<Input
			id="password-input"
			size="lg"
			placeholder="Введите ваш пароль"
			type="password"
			bind:value={valuePassword}
		/>
	</div>

	<div class="relative">
		<Button class="absolute right-0" on:click={auth}>Войти</Button>
	</div>
</div>

<svelte:window on:keydown={enter_press} />