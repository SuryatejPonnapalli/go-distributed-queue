package worker

import (
	"log"
	"strings"
	"time"

	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/common"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/llm"
	"github.com/SuryatejPonnapalli/go-distributed-queue/internal/queue"
)

func ProcessChatJob(job queue.ChatJob, svc *llm.LLMService) {
    log.Println("ChatJob started for:", job.Prompt)

    normalized := strings.ToLower(strings.TrimSpace(job.Prompt))
    key := "resp:" + normalized

    existingResp, err := common.Redis.Get(common.Ctx, key).Result()
    if err == nil && existingResp != "" {
        log.Println("Skipping ChatJob — response already exists for:", normalized)

        common.Redis.HSet(common.Ctx, "job:"+job.ID,
            "status", "done",
            "response", existingResp,
            "updated_at", time.Now().String(),
        )
        return
    }

    common.Redis.HSet(common.Ctx, "job:"+job.ID,
        "status", "chatting",
        "updated_at", time.Now().String(),
    )
    common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)

    resp, err := svc.GetPromptResponse(normalized)
    if err != nil {
        log.Println("chat failed:", err)

        common.Redis.HSet(common.Ctx, "job:"+job.ID,
            "status", "error",
            "error", err.Error(),
            "updated_at", time.Now().String(),
        )
        common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)
        return
    }

    common.Redis.Set(common.Ctx, key, resp, 7*24*time.Hour)

    log.Println("Stored response:", key)

    common.Redis.HSet(common.Ctx, "job:"+job.ID,
        "status", "done",
        "response", resp,
        "updated_at", time.Now().String(),
    )
    common.Redis.Expire(common.Ctx, "job:"+job.ID, 3*time.Hour)
}
