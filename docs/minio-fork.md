# Velero MinIO-aware build — `v1.18.2-minio`

**Image:** `docker.eqipe.cloud/velero/velero:v1.18.2-minio`

- Manifest list: `sha256:b8db8061…`
- `linux/amd64` `sha256:ab681ce4…`
- `linux/arm64` `sha256:e0207895…`

**Base:** upstream Velero `v1.18.2` (`c253c7fe…`) + one cherry-picked patch (`5316787`).

## What it fixes

On MinIO buckets with versioning **suspended** (the misleading `versioning: false`
default in the MinIO Helm chart), deleting a backup leaves an empty S3 prefix
behind. The S3 `list-objects` API still returns that prefix in `CommonPrefixes`,
so Velero's `ListBackups()` reports it as a valid backup — a "ghost" backup whose
metadata is gone but which Velero keeps trying to sync. This is upstream issue
[#8466](https://github.com/vmware-tanzu/velero/issues/8466) (reported by pdstefan,
still open).

## The patch

A guard in `pkg/persistence/object_store.go` `ListBackups()`: for each candidate
prefix, check that its `velero-backup.json` metadata object actually exists. If it
doesn't (or the existence check errors), skip the prefix instead of listing it as
a backup. Errors are logged so genuine storage problems stay visible; the benign
"metadata absent" case is silenced at debug level.

## Relationship to upstream

- **Issue #8466** — open, MinIO storage area, assigned to kaovilai & mpryc.
- **PR #9511** — open, implements the same `ListBackups()` guard; approved by
  chrisamti and kaovilai, pending shubham-pampattiwar and blackpiglet. Once merged
  and released upstream, the `-minio` fork can be retired.

## How to use

Point your Velero deployment's image at
`docker.eqipe.cloud/velero/velero:v1.18.2-minio`. Multi-arch, drop-in replacement
for upstream `v1.18.2`.

## Rebuilding for a new upstream release

The fork carries a single patch on top of an upstream release tag, using the
`vX.Y.Z-minio` naming convention for both branch and tag.

```bash
# 1. Fetch the upstream release tag
git fetch https://github.com/vmware-tanzu/velero.git \
  refs/tags/vX.Y.Z:refs/tags/vX.Y.Z

# 2. Branch from it and cherry-pick the MinIO fix
git checkout -b vX.Y.Z-minio vX.Y.Z
git cherry-pick 5316787

# 3. Tag and push (branch and tag share a name — use full refspecs)
git tag vX.Y.Z-minio
git push origin refs/heads/vX.Y.Z-minio:refs/heads/vX.Y.Z-minio
git push origin refs/tags/vX.Y.Z-minio:refs/tags/vX.Y.Z-minio
```

Then build the multi-arch image and push it to Harbor
(`docker.eqipe.cloud/velero/velero:vX.Y.Z-minio`).
