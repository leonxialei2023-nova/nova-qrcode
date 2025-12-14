Nova QR

Nova QR is a high-performance native PHP extension (.so) for generating QR codes, written in Go and exposed via a stable C ABI.

It is designed for server-side payment systems, high-concurrency QR generation, and low-latency environments, where pure PHP or user-space libraries are not sufficient.

---

Why Nova QR

Most QR code solutions in PHP are:
• Pure PHP (slow, CPU-heavy)
• ImageMagick dependent (heavy, unstable in containers)
• External services (latency, privacy, cost)

Nova QR solves this by moving QR generation into a native shared library.

Key Advantages
<br>• ⚡ Native performance via .so
<br>• 🧠 Low memory footprint
<br>• 🧩 No external dependencies
<br>• 🔒 Offline & privacy-safe
<br>• 🚀 Designed for high concurrency
<br>• 🧱 Clean C ABI for PHP extensions

---

Architecture

<br>└── PHP Extension (.so)
<br> └── C ABI Interface
<br> └── Go Core (QR generation)

    •	Core logic implemented in Go
    •	Exported as a shared library
    •	Loaded by PHP via a custom extension
    •	No shell calls, no external binaries

---

Features
• Generate QR codes from strings / URLs
• Output formats:
• PNG (default)
• Raw bitmap (optional)
• Customizable:
• Size
• Margin
• Error correction level
• Designed for:
• Alipay / WeChat Pay
• Payment deep links
• Order & transaction QR codes

---

Target Use Cases
• Payment systems
• Trading platforms
• High-frequency QR generation
• Headless / containerized PHP services
• ARM / x86 servers
• Embedded & edge devices

---

Status

🚧 Early development

Planned milestones:
• Core QR generator (Go)
• Stable C ABI
• PHP extension wrapper
• Benchmark & stress tests
• Documentation & examples

---

Installation (Planned)

```shell
# Build shared library
make build

# Install PHP extension
make install
```

PHP usage (planned):

```php
$png = nova_qr_encode("alipays://...");
file_put_contents("pay.png", $png);
```

Performance Goals
• ≥ 10x faster than pure PHP libraries
• Stable under high concurrency
• Predictable memory usage
• Zero runtime allocations in hot paths (where possible)

---

Design Principles
• Native first
• Minimal API surface
• Stable ABI
• Long-term maintainability
• No framework coupling

---

License

MIT License

---

About Nova

Nova is a long-term project focused on:
• Native performance
• Simple abstractions
• Cross-language interoperability
• Infrastructure-grade tooling

Nova QR is the first building block.
