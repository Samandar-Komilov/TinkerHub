# Chapter 9 — Modules and Packages (Working with Modules)

---

## 1. Importing third-party code

* You can pull external code directly from GitHub (or other VCS hosts).
* Example:

```go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("Generated ID:", id)
}
```

* Run:

  ```bash
  go mod init myapp
  go get github.com/google/uuid@latest
  go run .
  ```

👉 Go will download it, update `go.mod`, and store it in your module cache (`$GOPATH/pkg/mod`).

**Exercise 1:**
Use `github.com/fatih/color` to print colored text in terminal.
Initialize a module, import the package, and print `"Hello"` in red.

---

## 2. Working with versions

* Go modules follow **semantic versioning** (`vMAJOR.MINOR.PATCH`).
* You can pin specific versions:

```bash
go get github.com/google/uuid@v1.3.0
```

* Your `go.mod` will now list that version.

👉 If the author releases `v1.4.0`, you can upgrade safely (backward compatible).

**Exercise 2:**

* Install `github.com/google/uuid` at an older version (e.g., `v1.2.0`).
* Print its version from `go list -m all`.
* Then upgrade to the latest version and confirm it changed.

---

## 3. Minimum Version Selection (MVS)

* Go uses **minimum version selection**, not “latest wins.”
* If two dependencies require different versions of the same module:

  * Go picks the *minimum version* that still satisfies all constraints.
* This avoids “dependency hell” seen in npm, Maven, pip.

Example:

* Package A requires `uuid v1.1.0`.
* Package B requires `uuid v1.3.0`.
* Your project gets `uuid v1.3.0` (the minimum satisfying both).

**Exercise 3:**

* Create a module that imports **two libraries** that both depend on `google/uuid`.
* Run `go list -m all` and check which version Go picked.

---

## 4. Updating to compatible and incompatible versions

* **Compatible upgrade:** from `v1.2.0` → `v1.3.0`. Safe.
* **Incompatible upgrade:** Go enforces it by changing the **import path**.

  * `v2+` versions require `…/v2` in the import path.

Example:

```go
import "github.com/gin-gonic/gin"        // v1.x
import "github.com/gin-gonic/gin/v2"     // v2.x
```

👉 This avoids silent breaking changes.

**Exercise 4:**

* Find a package that has both `v1` and `v2` (e.g., `github.com/go-redis/redis/v8`).
* Try installing `v8` and notice you must import it as `/v8`.
* Compare how your `go.mod` changes.

---

## 5. Vendoring

* Vendoring = copy all dependencies into a `vendor/` folder.
* Run:

  ```bash
  go mod vendor
  ```
* Benefits:

  * Reproducible builds (no network needed).
  * Good for enterprise/government with restricted internet.
* Go will prefer `vendor/` if present.

**Exercise 5:**

* Create a module with `uuid`.
* Run `go mod vendor`.
* Delete your module cache (`$GOPATH/pkg/mod`) and run again → it will still work because vendored.

---

## 6. pkg.go.dev

* Official Go documentation site for all public modules.
* Example: [pkg.go.dev/github.com/google/uuid](https://pkg.go.dev/github.com/google/uuid)
* Shows docs, versions, import instructions.

**Exercise 6:**

* Visit `pkg.go.dev/github.com/fatih/color`.
* Look at functions, try one you didn’t know about (`color.Cyan`, `color.New`).

---

## 7. Versioning your module

* If you publish your own module:

  * Start with `v0.x` for experimental.
  * `v1.x` → stable API.
  * Breaking change → bump to `v2` and change import path (`/v2`).

Example:

* Suppose you write `github.com/ekzosfera/mathutils`.
* Release `v1.0.0`.
* Later, you break the API.
* Publish as `v2.0.0` and tell users to import:

  ```go
  import "github.com/ekzosfera/mathutils/v2"
  ```

**Exercise 7:**

* Create a local module `github.com/you/mathutils`.
* Publish (or simulate) `v1.0.0` with `func Add`.
* Make a breaking change (`func Add(a, b, c int)`) → bump to `/v2`.
* In another project, import both `v1` and `v2` in the same file.

---

# 🌟 Summary

* **Importing** → `go get` pulls from VCS.
* **Versions** → semantic versioning, pinned in `go.mod`.
* **MVS** → picks minimum version, avoids dependency hell.
* **Compatible vs incompatible upgrades** → `v1.x` vs `v2+` import paths.
* **Vendoring** → ship dependencies inside project.
* **pkg.go.dev** → docs & discovery.
* **Versioning your module** → follow semantic versioning and `/vN` rules.
