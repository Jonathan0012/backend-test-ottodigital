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

## Postman Collection

> https://www.postman.com/navigation-meteorologist-32062442/workspace/public-jrs/collection/23901344-cc74823a-f9a3-46e1-b499-0b5e074d5e4a?action=share&creator=23901344
