package main

/*

#include "../redismodule.h"

extern int GoOnLoad(RedisModuleCtx *ctx);

 int RedisModule_OnLoad(RedisModuleCtx *ctx) {

    if (RedisModule_Init(ctx, "MODULE_NAME", 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (GoOnLoad(ctx) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    // if (RedisModule_CreateCommand(
    //             ctx, "go.foo", GoCommand, "readonly", 1, 1, 1) == REDISMODULE_ERR) {
    //     return REDISMODULE_ERR;
    // }

    return REDISMODULE_OK;
}
*/
import "C"
