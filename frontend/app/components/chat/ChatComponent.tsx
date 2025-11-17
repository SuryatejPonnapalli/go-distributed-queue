import { useState } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Send } from "lucide-react";

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
    timestamp: "10:00 AM",
  },
  {
    id: "2",
    text: "Hi! I'd like to know more about your services.",
    sender: "user",
    timestamp: "10:01 AM",
  },
  {
    id: "3",
    text: "Of course! We offer a wide range of services including web development, mobile apps, and AI solutions. What interests you most?",
    sender: "bot",
    timestamp: "10:01 AM",
  },
  {
    id: "4",
    text: "Tell me about your web development services.",
    sender: "user",
    timestamp: "10:02 AM",
  },
  {
    id: "5",
    text: "We specialize in building modern, responsive web applications using React, Next.js, and Tailwind CSS. We can help you create everything from landing pages to complex full-stack applications.",
    sender: "bot",
    timestamp: "10:03 AM",
  },
  {
    id: "6",
    text: "That sounds great! Do you have any portfolio examples?",
    sender: "user",
    timestamp: "10:04 AM",
  },
  {
    id: "7",
    text: "Yes! We have built projects across various industries. I can share our portfolio with you. Would you like me to send you some examples?",
    sender: "bot",
    timestamp: "10:05 AM",
  },
];

export function ChatInterface() {
  const [messages, setMessages] = useState<Message[]>(SAMPLE_CONVERSATIONS);
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (!input.trim()) return;

    const userMessage: Message = {
      id: Date.now().toString(),
      text: input,
      sender: "user",
      timestamp: new Date().toLocaleTimeString([], {
        hour: "2-digit",
        minute: "2-digit",
      }),
    };

    setMessages([...messages, userMessage]);
    setInput("");

    // Simulate bot response
    setTimeout(() => {
      const botMessage: Message = {
        id: (Date.now() + 1).toString(),
        text:
          'Thanks for your message! I understand you said: "' +
          input +
          '". How else can I help?',
        sender: "bot",
        timestamp: new Date().toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
        }),
      };
      setMessages((prev) => [...prev, botMessage]);
    }, 500);
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
