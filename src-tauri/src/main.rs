#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tauri::{AppHandle, LogicalSize, Manager, Size};

fn main() {
    tauri::Builder::default()
        .invoke_handler(tauri::generate_handler![href_auth, href_main, auth])
        .run(tauri::generate_context!())
        .expect("error: run()");
}

#[tauri::command]
fn href_auth(app: AppHandle) {
    if app.get_window("main").is_none() {
        return ();
    }

    let main = app.get_window("main").unwrap();
    main.set_title("Авторизация").expect("error: set_title()");
    main.set_resizable(false).expect("error: set_resizable()");
    main.set_size(Size::Logical(LogicalSize {
        width: 400.0_f64,
        height: 300.0_f64,
    }))
    .expect("error: set_size()");
    main.center().expect("error: center()");
}

#[tauri::command]
fn href_main(app: AppHandle) {
    if app.get_window("main").is_none() {
        return ();
    }

    let main = app.get_window("main").unwrap();
    main.set_title("Информаионная система тира")
        .expect("error: set_title()");
    main.set_resizable(true).expect("error: set_resizable()");
    main.set_size(Size::Logical(LogicalSize {
        width: 1280.0_f64,
        height: 720.0_f64,
    }))
    .expect("error: set_size()");
    main.center().expect("error: center()");
}

#[tauri::command]
fn auth(login: String, password: String) -> bool {
    if login == "login" && password == "password" {
        return true;
    } else {
        return false;
    }
}
