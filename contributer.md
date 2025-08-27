# Contributing to Maven Dependency Updater

Thanks for your interest in contributing! 🚀
This project is still early-stage, and contributions of all kinds are welcome — from bug reports to new features.

## How to Contribute

### 1. Fork & Clone

* Fork the repo: [Maven Package Manager](https://github.com/iamritikbhardwaj/maven-package-manager)
* Clone your fork:

  ```bash
  git clone https://github.com/<your-username>/maven-package-manager.git
  cd maven-package-manager
  ```

### 2. Create a Branch

Use a descriptive branch name:

```bash
git checkout -b feature/add-maven-central-support
```

### 3. Make Your Changes

* Keep code clean and consistent (Go formatting: `go fmt ./...`)
* Write clear commit messages
* Add comments where needed

### 4. Run & Test

Make sure the tool still works as expected:

```bash
go run main.go <groupId> <artifactId> <version>
```

### 5. Push & Open PR

```bash
git push origin feature/add-maven-central-support
```

Then open a **Pull Request** against the `main` branch.

## Contribution Ideas

Here are some open enhancements you could work on:

* [ ] Support creating `<dependencies>` if missing
* [ ] Fetch latest versions from Maven Central API
* [ ] Improve XML formatting with a linter
* [ ] Add proper error handling & validation
* [ ] Unit tests for different pom.xml structures

## Code of Conduct

Be respectful, constructive, and collaborative. The goal is to learn, share, and build together.

## License

By contributing, you agree that your code will be licensed under the MIT License.
