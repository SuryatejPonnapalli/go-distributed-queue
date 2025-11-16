from dotenv import load_dotenv
load_dotenv()

from fastapi import FastAPI
from models import embed_request, llm_service
from embedding import get_embedding
from chat import ask_ollama


app = FastAPI()
EmbedRequest = embed_request.EmbedRequest
ChatRequest = llm_service.ChatRequest

@app.post("/embed")
def embed(req: EmbedRequest):
    vector = get_embedding(req.text)
    return {"embedding": vector}

@app.post("/chat")
def chat(req: ChatRequest):
    response = ask_ollama(req.text)
    return {"response": response}