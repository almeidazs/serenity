<div align="center">

# Serenity

<p>
  <strong>An aggressive, no-noise and ultra fast Go linter.</strong>
</p>

<p>
  Serenity is a configurable Go code quality tool built around focused AST rules,
  fast package traversal, deterministic issue limits, inline source frames, cache
  support and automatic fixes where the rewrite is safe.
</p>

<p>
  <a href="#installation">Installation</a> ·
  <a href="#quick-start">Quick Start</a> ·
  <a href="#usage">Usage</a> ·
  <a href="#configuration">Configuration</a> ·
  <a href="#rules">Rules</a> ·
  <a href="#suppression">Suppression</a> ·
  <a href="#contributing">Contributing</a>
</p>

</div>

<div align="center">

<h2 id="why-serenity">Why Serenity?</h2>

<p>
  Serenity is designed for Go projects that want sharp defaults without a noisy
  feedback loop. It reports issues with enough context to fix them immediately,
  limits output when asked, and can apply selected rewrites directly to source
  files.
</p>

<table align="center">
  <tr>
    <th>Area</th>
    <th>What Serenity does</th>
  </tr>
  <tr>
    <td><strong>Signal</strong></td>
    <td>Runs explicit rule groups with severity levels and max issue limits.</td>
  </tr>
  <tr>
    <td><strong>Speed</strong></td>
    <td>Uses concurrent package analysis, optional cache and deterministic single-worker mode when limits are enabled.</td>
  </tr>
  <tr>
    <td><strong>Fixes</strong></td>
    <td>Marks fixable diagnostics and supports safe <code>--write</code> rewrites plus explicit <code>--unsafe</code> rewrites.</td>
  </tr>
  <tr>
    <td><strong>Config</strong></td>
    <td>Loads JSON, YAML or TOML from the current tree, an explicit path, or <code>SERENITY_CONFIG_PATH</code>.</td>
  </tr>
  <tr>
    <td><strong>Output</strong></td>
    <td>Prints severity, rule name, numeric rule id, source location, a compact code frame and a human hint.</td>
  </tr>
</table>

<h2 id="installation">Installation</h2>

<p>
  Install the latest version with the Go toolchain:
</p>

</div>

```sh
go install github.com/serenitysz/serenity@latest
```

<div align="center">

<p>
  Or build it directly from the repository:
</p>

</div>

```sh
git clone https://github.com/serenitysz/serenity.git
cd serenity
go build -o serenity .
```

<div align="center">

<p>
  Release archives are produced with GoReleaser for Linux, macOS and Windows on
  <code>amd64</code> and <code>arm64</code>.
</p>

<h2 id="quick-start">Quick Start</h2>

</div>

```sh
# Create serenity.json with the recommended preset.
serenity init

# Check the current module.
serenity check ./...

# Apply safe automatic fixes.
serenity check --write ./...

# Apply safe fixes and explicitly allowed unsafe fixes.
serenity check --write --unsafe ./...
```

<div align="center">

<p>
  For an interactive setup that can choose JSON, YAML or TOML, strict mode and
  autofix defaults, run:
</p>

</div>

```sh
serenity init --interactive
```

<div align="center">

<h2 id="usage">Usage</h2>

<table align="center">
  <tr>
    <th>Command</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>serenity check [path...]</code></td>
    <td>Analyze one or more files or directories. Defaults to the current directory.</td>
  </tr>
  <tr>
    <td><code>serenity init</code></td>
    <td>Create a default <code>serenity.json</code> configuration.</td>
  </tr>
  <tr>
    <td><code>serenity init -i</code></td>
    <td>Interactively choose config format, path, strict preset and autofix.</td>
  </tr>
  <tr>
    <td><code>serenity status</code></td>
    <td>Print version, commit, platform and discovered config status.</td>
  </tr>
  <tr>
    <td><code>serenity docs</code></td>
    <td>Open the Serenity documentation site in the default browser.</td>
  </tr>
  <tr>
    <td><code>serenity update</code></td>
    <td>Self-update the installed binary from GitHub releases.</td>
  </tr>
</table>

<h3>Common flags</h3>

<table align="center">
  <tr>
    <th>Flag</th>
    <th>Scope</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><code>--config path</code></td>
    <td>global/check</td>
    <td>Load a specific config file instead of auto-discovery.</td>
  </tr>
  <tr>
    <td><code>--no-color</code></td>
    <td>global</td>
    <td>Disable ANSI colors in output.</td>
  </tr>
  <tr>
    <td><code>--verbose</code></td>
    <td>global</td>
    <td>Enable additional diagnostics where supported.</td>
  </tr>
  <tr>
    <td><code>--max-issues n</code></td>
    <td>check</td>
    <td>Stop after <code>n</code> issues. <code>0</code> means unlimited.</td>
  </tr>
  <tr>
    <td><code>--max-file-size bytes</code></td>
    <td>check</td>
    <td>Skip files larger than the provided byte limit.</td>
  </tr>
  <tr>
    <td><code>--write</code></td>
    <td>check</td>
    <td>Apply safe automatic fixes and format changed files.</td>
  </tr>
  <tr>
    <td><code>--unsafe</code></td>
    <td>check</td>
    <td>Allow unsafe fixes. Must be used with <code>--write</code>.</td>
  </tr>
</table>

<h2 id="configuration">Configuration</h2>

<p>
  Serenity auto-discovers one of these files from the current working directory
  upward:
</p>

<p>
  <code>serenity.json</code> · <code>serenity.yaml</code> · <code>serenity.yml</code> · <code>serenity.toml</code>
</p>

<p>
  Discovery order can be bypassed with <code>--config</code> or with the
  <code>SERENITY_CONFIG_PATH</code> environment variable. If no config is found,
  Serenity runs with the generated default configuration and applies the
  recommended preset.
</p>

<h3>Minimal JSON config</h3>

</div>

```json
{
	"$schema": "https://raw.githubusercontent.com/serenitysz/schema/main/versions/1.0.0.json",
	"linter": {
		"use": true,
		"rules": {
			"recommended": true
		},
		"issues": {
			"use": true,
			"max": 20
		}
	},
	"assistance": {
		"use": true,
		"autofix": false
	},
	"performance": {
		"use": true,
		"caching": true
	}
}
```

<div align="center">

<h3>Rule group config</h3>

<p>
  Rules are grouped by domain. Each enabled group uses <code>"use": true</code>,
  and each rule accepts a <code>severity</code> of <code>error</code>,
  <code>warn</code> or <code>info</code>. Rules with limits or patterns expose
  extra fields such as <code>max</code>, <code>pattern</code>,
  <code>packages</code>, <code>maxSize</code> or
  <code>maxUnnamedSameType</code>.
</p>

</div>

```json
{
	"linter": {
		"use": true,
		"rules": {
			"recommended": false,
			"imports": {
				"use": true,
				"noDotImports": {
					"severity": "error"
				},
				"disallowedPackages": {
					"severity": "warn",
					"packages": ["log", "log/slog"]
				},
				"redundantImportAlias": {
					"severity": "warn"
				}
			},
			"bestPractices": {
				"use": true,
				"useContextInFirstParam": {
					"severity": "warn"
				},
				"maxParams": {
					"severity": "warn",
					"max": 4
				},
				"useSliceCapacity": {
					"severity": "warn"
				}
			},
			"complexity": {
				"use": true,
				"maxFuncLines": {
					"severity": "warn",
					"max": 35
				},
				"maxLineLength": {
					"severity": "warn",
					"max": 100
				}
			}
		},
		"issues": {
			"use": true,
			"max": 100
		}
	},
	"assistance": {
		"use": true,
		"autofix": false
	},
	"performance": {
		"use": true,
		"threads": 8,
		"caching": true
	}
}
```

<div align="center">

<h3>Presets</h3>

<table align="center">
  <tr>
    <th>Preset</th>
    <th>How to enable</th>
    <th>Behavior</th>
  </tr>
  <tr>
    <td><strong>Recommended</strong></td>
    <td><code>"recommended": true</code> or <code>serenity init</code></td>
    <td>Turns on a balanced set of import, best-practice and complexity rules with a default max issue count.</td>
  </tr>
  <tr>
    <td><strong>Strict</strong></td>
    <td><code>serenity init --interactive</code> and answer yes to strict mode</td>
    <td>Creates an explicit config with stricter severities and lower limits for teams that want tighter gates.</td>
  </tr>
</table>

<h3>Performance config</h3>

<p>
  <code>performance.caching</code> enables the local lint cache. The cache is
  keyed by Serenity version, cache version, mutating mode, unsafe mode and the
  loaded configuration. <code>performance.threads</code> overrides the default
  worker count, which otherwise follows <code>GOMAXPROCS</code>.
</p>

<h3>Assistance config</h3>

<p>
  <code>assistance.autofix</code> allows safe fixes to run without passing
  <code>--write</code>. Unsafe fixes still require <code>--write --unsafe</code>.
</p>

<h2 id="rules">Rules</h2>

<p>
  The table below lists the rules wired into the current runner. The names match
  CLI output and suppression comments.
</p>

<table align="center">
  <tr>
    <th>Group</th>
    <th>Config key</th>
    <th>Rule name</th>
    <th>Fix support</th>
  </tr>
  <tr>
    <td>Errors</td>
    <td><code>errorStringFormat</code></td>
    <td><code>error-string-format</code></td>
    <td>Safe</td>
  </tr>
  <tr>
    <td>Errors</td>
    <td><code>errorNotWrapped</code></td>
    <td><code>error-not-wrapped</code></td>
    <td>Safe</td>
  </tr>
  <tr>
    <td>Imports</td>
    <td><code>noDotImports</code></td>
    <td><code>no-dot-imports</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Imports</td>
    <td><code>disallowedPackages</code></td>
    <td><code>disallowed-packages</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Imports</td>
    <td><code>redundantImportAlias</code></td>
    <td><code>redundant-import-alias</code></td>
    <td>Safe</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>maxParams</code></td>
    <td><code>max-params</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>useContextInFirstParam</code></td>
    <td><code>context-first-param</code></td>
    <td>Unsafe</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>avoidEmptyStructs</code></td>
    <td><code>avoid-empty-structs</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>noMagicNumbers</code></td>
    <td><code>no-magic-numbers</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>alwaysPreferConst</code></td>
    <td><code>always-prefer-const</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>noDeferInLoop</code></td>
    <td><code>no-defer-in-loop</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>useSliceCapacity</code></td>
    <td><code>use-slice-capacity</code></td>
    <td>Safe</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>noBareReturns</code></td>
    <td><code>no-bare-returns</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Best practices</td>
    <td><code>getMustReturnValue</code></td>
    <td><code>get-must-return-value</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Correctness</td>
    <td><code>emptyBlock</code></td>
    <td><code>empty-block</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Correctness</td>
    <td><code>ambiguousReturns</code></td>
    <td><code>ambiguous-return</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Correctness</td>
    <td><code>boolLiteralExpressions</code></td>
    <td><code>boolean-literal-expressions</code></td>
    <td>Safe</td>
  </tr>
  <tr>
    <td>Complexity</td>
    <td><code>maxFuncLines</code></td>
    <td><code>max-func-lines</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Complexity</td>
    <td><code>maxLineLength</code></td>
    <td><code>max-line-length</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Naming</td>
    <td><code>receiverNames</code></td>
    <td><code>receiver-name</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Naming</td>
    <td><code>exportedIdentifiers</code></td>
    <td><code>exported-identifiers</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Naming</td>
    <td><code>importedIdentifiers</code></td>
    <td><code>imported-identifiers</code></td>
    <td>No</td>
  </tr>
  <tr>
    <td>Style</td>
    <td><code>preferIncDec</code></td>
    <td><code>prefer-inc-dec</code></td>
    <td>Safe</td>
  </tr>
</table>

<h2 id="autofix">Autofix</h2>

<p>
  Serenity separates safe and unsafe rewrites. Safe rewrites are applied with
  <code>--write</code> or by setting <code>assistance.autofix</code> to
  <code>true</code>. Unsafe rewrites require both <code>--write</code> and
  <code>--unsafe</code>.
</p>

<table align="center">
  <tr>
    <th>Mode</th>
    <th>Command</th>
    <th>Effect</th>
  </tr>
  <tr>
    <td>Report only</td>
    <td><code>serenity check</code></td>
    <td>Shows diagnostics and marks fixable issues.</td>
  </tr>
  <tr>
    <td>Safe rewrite</td>
    <td><code>serenity check --write</code></td>
    <td>Applies safe fixes and runs Go formatting on changed files.</td>
  </tr>
  <tr>
    <td>Unsafe rewrite</td>
    <td><code>serenity check --write --unsafe</code></td>
    <td>Also applies fixes that can change public signatures or call contracts.</td>
  </tr>
</table>

<h2 id="suppression">Suppression</h2>

<p>
  Suppressions use rule names, not config keys. Inline suppressions apply to the
  comment line and the following line. File-wide suppressions must appear before
  the package declaration. Unused suppressions are reported as warnings.
</p>

</div>

```go
// @serenity-ignore no-magic-numbers: port is part of the protocol
const defaultPort = 8080

// @serenity-ignore-all no-dot-imports: generated test fixture
package fixture
```

<div align="center">

<h2 id="ci">CI</h2>

<p>
  Serenity works well as a focused CI gate. A minimal GitHub Actions job can
  install the CLI and run it against the repository:
</p>

</div>

```yaml
name: Serenity

on:
  pull_request:
  push:

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: stable
      - run: go install github.com/serenitysz/serenity@latest
      - run: serenity check ./...
```

<div align="center">

<h2 id="exit-codes">Exit Codes</h2>

<table align="center">
  <tr>
    <th>Code</th>
    <th>Meaning</th>
  </tr>
  <tr>
    <td><code>0</code></td>
    <td>No issues or command completed successfully.</td>
  </tr>
  <tr>
    <td><code>1</code></td>
    <td>Command-level failure, including lint issues.</td>
  </tr>
  <tr>
    <td><code>2</code></td>
    <td>Internal error such as config, parse or filesystem failures.</td>
  </tr>
</table>

<h2 id="development">Development</h2>

</div>

```sh
go test ./...
gofmt -w .
go run . check ./...
```

<div align="center">

<p>
  The repository uses Go tests for the linter, config loading, output rendering,
  exception handling and rule messages. Release artifacts are built through the
  release workflow and GoReleaser configuration in this repository.
</p>

<h2 id="contributing">Contributing</h2>

<p>
  Contributions are welcome. Please read
  <a href="CONTRIBUTING.md">CONTRIBUTING.md</a> before opening a pull request,
  and use the issue templates for bug reports or feature requests.
</p>

<p>
  Security reports should follow <a href="SECURITY.md">SECURITY.md</a>.
</p>

<h2 id="license">License</h2>

<p>
  Serenity is released under the <a href="LICENSE">MIT License</a>.
</p>

</div>
