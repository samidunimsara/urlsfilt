``` # Filter out URLs containing "www."
./urlsfilt -i urls.txt -f www.

# Filter multiple patterns
./urlsfilt -i urls.txt -f www.,.js,outfit

# Save output to a file
./urlsfilt -i urls.txt -f www.,outfit -o filtered.txt
```

## How it works

Based on your example input:
- `urlsfilt -i urls.txt -f www.` → Excludes `https://www.stylehint.com/...` URLs
- `urlsfilt -i urls.txt -f www.,.js,outfit` → Excludes URLs with "www.", ".js", or "outfit"

**Output would be:**
```
https://sss.stylehint.com/jp/ja/outfit/
```
(if you filter by `www.` only)

Or just:
```
https://eee.stylehint.com/jp/ja/outfit/1.91.0
