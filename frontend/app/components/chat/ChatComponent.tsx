import { useState } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Send } from "lucide-react";
import { postRequest } from "~/utils/api/post";
import { getRequest } from "~/utils/api/get";

interface Message {
  id: string;
  text: string;
  sender: "user" | "bot";
  timestamp: string;
}

const SAMPLE_CONVERSATIONS: Message[] = [
  {
    id: "1",
    text: "Hello! How can I help you today?",
    sender: "bot",
    timestamp: new Date().toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
    }),
  },
];

export function ChatInterface() {
  const [messages, setMessages] = useState<Message[]>(SAMPLE_CONVERSATIONS);
  const [input, setInput] = useState("");
  const [loading, setLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");

  const handleSend = async () => {
    if (!input.trim()) {
      setErrorMsg("Enter a prompt");
      return;
    }

    const userMessage: Message = {
      id: Date.now().toString(),
      text: input,
      sender: "user",
      timestamp: new Date().toLocaleTimeString([], {
        hour: "2-digit",
        minute: "2-digit",
      }),
    };

    setMessages((prev) => [...prev, userMessage]);
    setInput("");
    setLoading(true);

    try {
      const res = await postRequest("llm/embed", { prompt: input });

      if (res.status !== 200) {
        setErrorMsg("Failed to send prompt");
        setLoading(false);
        return;
      }

      if (
        res.data.status == "cached_exact" ||
        res.data.status == "semantic_reuse"
      ) {
        const botMessage: Message = {
          id: (Date.now() + 1).toString(),
          text: res.data.response,
          sender: "bot",
          timestamp: new Date().toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          }),
        };
        setMessages((prev) => [...prev, botMessage]);
        return;
      }

      const jobID = res.data.jobID;

      // -----------------------------
      // POLLING LOOP WITH TIMEOUT
      // -----------------------------
      const timeoutMs = 50000; // 20 sec timeout
      const intervalMs = 1000; // poll every 1 sec
      const start = Date.now();

      let jobData = null;

      while (Date.now() - start < timeoutMs) {
        const poll = await getRequest(`llm/jobs/${jobID}`);
        jobData = poll.data;
        console.log("job", jobData);

        if (jobData.data.status === "done") break;
        if (jobData.data.status === "error") break;

        // still embedding or chatting
        await new Promise((r) => setTimeout(r, intervalMs));
      }

      setLoading(false);

      if (!jobData) {
        setErrorMsg("Request timed out. Try again.");
        return;
      }

      if (jobData.status === "error") {
        setErrorMsg(jobData.error || "Something went wrong");
        return;
      }

      if (jobData.data.status === "done") {
        const botMessage: Message = {
          id: (Date.now() + 1).toString(),
          text: jobData.data.response,
          sender: "bot",
          timestamp: new Date().toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          }),
        };

        setMessages((prev) => [...prev, botMessage]);
      }
    } catch (err) {
      setErrorMsg("Network error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full max-w-2xl h-screen md:h-[600px] bg-card rounded-lg border border-border shadow-lg flex flex-col">
      {/* Header */}
      <div className="border-b border-border p-4 bg-primary text-primary-foreground rounded-t-lg">
        <h2 className="text-xl font-bold">Support Assistant</h2>
        <p className="text-sm opacity-90">Always here to help</p>
      </div>

      {/* Messages Container */}
      <div className="flex-1 overflow-y-auto p-4 space-y-4">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`flex ${
              message.sender === "user" ? "justify-end" : "justify-start"
            }`}
          >
            <div
              className={`max-w-xs lg:max-w-md px-4 py-2 rounded-lg ${
                message.sender === "user"
                  ? "bg-primary text-primary-foreground rounded-br-none"
                  : "bg-muted text-muted-foreground rounded-bl-none"
              }`}
            >
              <p className="text-sm">{message.text}</p>
              <span
                className={`text-xs mt-1 block ${
                  message.sender === "user" ? "opacity-75" : "opacity-60"
                }`}
              >
                {message.timestamp}
              </span>
            </div>
          </div>
        ))}
      </div>

      {/* Input Area */}
      <div className="border-t border-border p-4 bg-card rounded-b-lg">
        <div className="flex gap-2">
          <Input
            placeholder="Type your message..."
            value={input}
            onChange={(e) => setInput(e.target.value)}
            onKeyPress={(e) => {
              if (e.key === "Enter") handleSend();
            }}
            className="flex-1"
          />
          <Button
            onClick={handleSend}
            disabled={!input.trim()}
            size="icon"
            className="bg-primary hover:bg-primary/90"
          >
            <Send className="w-4 h-4" />
          </Button>
        </div>
      </div>
    </div>
  );
}
