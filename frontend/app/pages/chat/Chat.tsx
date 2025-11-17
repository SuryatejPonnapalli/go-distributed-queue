import { ChatInterface } from "~/components/chat/ChatComponent";

export const metadata = {
  title: "Chatbot",
  description: "AI Chatbot Interface",
};

export default function ChatPage() {
  return (
    <main className="min-h-screen bg-background flex items-center justify-center p-4">
      <ChatInterface />
    </main>
  );
}
