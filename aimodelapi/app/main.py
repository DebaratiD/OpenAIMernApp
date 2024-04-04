from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from huggingface_hub import InferenceClient
import io
import base64


HF_TOKEN ="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"
inference = InferenceClient(model="runwayml/stable-diffusion-v1-5", token=HF_TOKEN)

class PromptModel(BaseModel):
    prompt:str

app = FastAPI()
origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


# model_id = "runwayml/stable-diffusion-v1-5"
# pipe = StableDiffusionPipeline.from_pretrained(model_id, torch_dtype=torch.float16)
# pipe = pipe.to("cuda")


@app.get("/")
def root():
    return {"message":"Hello from stable diffusion!"}

@app.post("/create")
def CreateImage(req:PromptModel):
    image = inference.text_to_image(req.prompt)
    data={}
    buffered = io.BytesIO()
    image.save(buffered, format="JPEG")
    img_str = base64.b64encode(buffered.getvalue())
    #img = cv2.imencode('.jpg',image)
    # convert the image to json
    data['photo'] = img_str

    
    return data