# GOLANG REST API USER TASK
---
## About the project
This is a simple REST Api that uses Go language and Gin Gonic Framework

## Tables

a. **users**
- `id` (integer, primary key, auto increment)
- `name` (string, max 255 characters)
- `email` (string, max 255 characters, unique)
- `password` (string, max 255 characters)
- `created_at` (datetime, with current timestamp)
- `updated_at` (datetime, with current timestamp on update)

b. **tasks**
- `id` (integer, primary key, auto increment)
- `user_id` (foreign key to users)
- `title` (string, max 255 characters)
- `description` (text)
- `status` (string, max 50 characters, default 'pending')
- `created_at` (datetime, with current timestamp)
- `updated_at` (datetime, with current timestamp on update)

## Endpoints


- Endpoint `/users`
    - `POST`: Register user baru.
    - `GET`: Dapatkan daftar semua pengguna.

- Endpoint `/users/:id`
    - `GET`: Dapatkan detail pengguna berdasarkan `id`.
    - `PUT`: Update informasi pengguna berdasarkan `id`.
    - `DELETE`: Hapus pengguna berdasarkan `id`.

- Endpoint `/tasks`
    - `POST`: Membuat task baru.
    - `GET`: Dapatkan daftar semua task.

- Endpoint `/tasks/:id`
    - `GET`: Dapatkan detail task berdasarkan `id`.
    - `PUT`: Update informasi task berdasarkan `id`.
    - `DELETE`: Hapus task berdasarkan `id`.

