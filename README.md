# 🔐 Ed25519 Key Generator

This is a simple Go-based service that generates an Ed25519 key pair (private and public) and outputs them in base64-encoded format. It can be built and run using Docker.

## 📦 Table of Contents

- [Features](#features)
- [Key Format](#key-format)
- [Build & Run](#build--run)
- [Usage](#usage)
- [Example Output](#example-output)

---

## 🛠 Features

The program performs the following:

1. Generates a 32-byte cryptographically secure random seed.
2. Derives an Ed25519 key pair from the seed.
3. Encodes:
    - The **private key** as a 64-byte value: \`[seed || publicKey]\`
    - The **public key** as a 32-byte value.
4. Outputs both values in base64 URL-safe encoding.

---

## 🔐 Key Format

- **PRIVATE** – base64 URL-encoded 64-byte value (\`seed + public key\`)
- **PUBLIC** – base64 URL-encoded 32-byte public key

These formats are compatible with many cryptographic systems expecting raw Ed25519 keys.

---

## 🐳 Build & Run

### 📁 Project Structure

```
.
├── Dockerfile
├── main.go
└── README.md
```

### 🔨 Build the Docker Image

```bash
docker build -t ed25519-gen .
```

### ▶️ Run the Service

```bash
docker run --rm ed25519-gen
```

---

## 🧪 Usage

You can run the binary directly:

```bash
./genkeys
```

Or use Docker:

```bash
docker run --rm ed25519-gen
```

---

## 🧾 Example Output

```
PRIVATE=wO5u6DfyAW...<truncated>...X5g
PUBLIC=Py3iNUJXtF...<truncated>...Eg
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
