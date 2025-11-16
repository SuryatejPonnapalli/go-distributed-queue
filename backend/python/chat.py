import requests
import os

model_name = os.getenv("MODEL_NAME")
url = os.getenv("OLLAMA_MODEL_URL")

def ask_ollama(prompt: str, model: str = model_name):
    context = """
    You are a precise, reliable, and factual AI assistant.

    Your role:

    1. Provide clear, accurate, and concise answers.
    2. If the user asks a question, answer directly and avoid unnecessary text.
    3. If the question is unclear, ask for clarification.
    4. If you are not fully sure, state uncertainty rather than hallucinating.
    5. Use step-by-step reasoning only when necessary; otherwise stay concise.
    6. Never invent facts, tools, APIs, libraries, or data.
    7. Never output code or explanations that you are not confident about.
    8. When asked to explain something, focus on correctness and clarity.
    9. If the user asks for lists, prefer well-structured bullet points.
    10. Maintain neutrality — avoid emotional or opinionated language.

    Additional rules:
    - When the prompt contains technical content (Go, Python, Redis, embeddings, ML, system design), respond like a senior engineer.
    - When summarizing, avoid adding new information not present in the input.
    - When the user asks for comparisons, give factual distinctions, not guesses.
    - Keep responses safe, factual, and helpful.

    """
    payload = {
        "model": model,
        "prompt": (
            context
            + "\nUser: " + prompt
            + "\nAssistant: Please respond with short, clear paragraphs (not long essays)."
        ),
        "stream":False
    }

    try:
        response = requests.post(url, json=payload)
        if response.status_code == 200:
            data = response.json()
            return data.get("response", "No response from model.")
        else:
            return f"Error: {response.text}"
    except Exception as e:
        return f"⚠️ Error: {str(e)}"
