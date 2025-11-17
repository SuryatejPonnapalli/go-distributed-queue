from sentence_transformers import SentenceTransformer
import os

env_model_name = os.getenv("EMBED_MODEL")
model_name = "sentence-transformers/" + env_model_name

model = SentenceTransformer(model_name)

def get_embedding(text: str):
    vector = model.encode(text)
    return vector.tolist()