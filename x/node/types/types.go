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

const TEXT_NODE_STATUS_NA = "Inactive"
const TEXT_NODE_STATUS_ONLINE = "Active"
const TEXT_NODE_STATUS_SERVE_GATEWAY = "Gateway"
const TEXT_NODE_STATUS_SERVE_STORAGE = "Storage Provider"
const TEXT_NODE_STATUS_ACCEPT_ORDER = "Storage Service"
const TEXT_NODE_STATUS_SERVE_INDEXING = "Indexing"
const TEXT_NODE_STATUS_SERVE_FISHING = "Fishing"
const TEXT_NODE_STATUS_SERVE_JAILED = "Jailed"

const NODE_SUBSIDY_HALVING_INTERVAL int64 = 2 * 365 * 24 * 60 * 60 / 5

const FISHMEN_LIST_DEPOSITOR = "sao1523r6gz79f2w46nx26q4x5vylqgw6pwf7uem9c"
const NODE_NORMAL = 0
const NODE_SUPER = 1

const TEXT_NODE_NORMAL = "Normal Node"
const TEXT_NODE_SUPER = "Super Node"
const NODE_STATUS_SUPER_REQUIREMENT uint32 = NODE_STATUS_ONLINE | NODE_STATUS_SERVE_GATEWAY | NODE_STATUS_SERVE_STORAGE | NODE_STATUS_ACCEPT_ORDER
