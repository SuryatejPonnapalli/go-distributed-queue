from fastapi import FastAPI
from models import embed_request
from embedding import get_embedding

app = FastAPI()
EmbedRequest = embed_request.EmbedRequest

@app.post("/embed")
def embed(req: EmbedRequest):
    vector = get_embedding(req.text)
    return {"embedding": vector}