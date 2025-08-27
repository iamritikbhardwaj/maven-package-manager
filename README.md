# 🛠️ Maven Dependency Updater (Go CLI Tool)

A simple CLI tool written in **Go** to add or update dependencies in a Maven `pom.xml` file — similar to `npm install`.

---

## 🚀 Features

- ✅ Adds new dependencies to your `pom.xml`
- ✅ Updates the version if the dependency already exists
- ✅ Keeps XML structure clean and well-formatted
- ✅ Minimal dependencies, cross-platform

---

## 📂 Project Structure

```

maven-dep-updater/
├── main.go
└── pom.xml         <-- Your Maven project's pom.xml

````

---

## ⚙️ Requirements

- [Go](https://golang.org/dl/) 1.18 or higher
- A valid `pom.xml` file in the same directory

---

## 📥 Installation

### Option 1: Run Directly

```bash
go run main.go <groupId> <artifactId> <version>
````

### Example

```bash
go run main.go org.apache.commons commons-lang3 3.12.0
```

### Option 2: Build as Binary

```bash
go build -o maven-add-dep main.go
./maven-add-dep org.springframework spring-context 6.2.11
```

---

## 🧠 How It Works

- The tool parses `pom.xml` using Go's `encoding/xml`
- It searches for an existing dependency by `groupId` and `artifactId`
- If found:

  - The `<version>` is updated
- If not:

  - A new `<dependency>` block is added under `<dependencies>`
- The updated `pom.xml` is written back

---

## 📝 Sample Output

Before:

```xml
<dependencies>
  <dependency>
    <groupId>org.apache.commons</groupId>
    <artifactId>commons-lang3</artifactId>
    <version>3.10</version>
  </dependency>
</dependencies>
```

After:

```xml
<dependencies>
  <dependency>
    <groupId>org.apache.commons</groupId>
    <artifactId>commons-lang3</artifactId>
    <version>3.12.0</version>
  </dependency>
  <dependency>
    <groupId>org.springframework</groupId>
    <artifactId>spring-context</artifactId>
    <version>6.2.11</version>
  </dependency>
</dependencies>
```

---

## ⚠️ Notes

- The script assumes the `pom.xml` structure is standard (with a `<dependencies>` section).
- It doesn’t resolve transitive dependencies — it only updates the XML.
- Always commit or back up your `pom.xml` before running.

---

## 📌 TODO (Optional Enhancements)

- [ ] Support inserting `<dependencies>` if missing
- [ ] Format output using `go/format` or an XML linter
- [ ] Fetch latest versions from Maven Central (via API)
- [ ] Validate input format

---

## 📄 License

MIT License

---

## 👨‍💻 Author

Created by \[Ritik Singh]
