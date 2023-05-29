package types

type NodeStatus uint32

const NODE_STATUS_NA uint32 = 0
const NODE_STATUS_ONLINE uint32 = 1
const NODE_STATUS_SERVE_GATEWAY uint32 = 1 << 1
const NODE_STATUS_SERVE_STORAGE uint32 = 1 << 2
const NODE_STATUS_ACCEPT_ORDER uint32 = 1 << 3
const NODE_STATUS_SERVE_INDEXING uint32 = 1 << 4
const NODE_STATUS_SERVE_FISHING uint32 = 1 << 5

const NODE_STATUS_SERVE_JAILED uint32 = 1 << 11

const NODE_SUBSIDY_HALVING_INTERVAL int64 = 2 * 365 * 24 * 60 * 60 / 5
