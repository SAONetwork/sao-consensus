package types

type NodeStatus uint32

const NODE_STATUS_NA uint32 = 0
const NODE_STATUS_ONLINE uint32 = 1
const NODE_STATUS_SERVE_GATEWAY uint32 = 1 << 1
const NODE_STATUS_SERVE_STORAGE uint32 = 1 << 2
const NODE_STATUS_ACCEPT_ORDER uint32 = 1 << 3
const NODE_SUBSIDY_HALVING_INTERVAL int64 = 2 * 365 * 24 * 60 * 60 / 5

// super node share threshold
const SHARE_THRESHOLD int64 = 10

const NODE_NORMAL = 0
const NODE_SUPER = 1
