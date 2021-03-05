from fastapi import FastAPI, Request, Form, File, UploadFile
from enum import Enum
from pydantic import BaseModel
from typing import Optional
from fastapi.templating import Jinja2Templates
from fastapi.staticfiles import StaticFiles

app = FastAPI()

app.mount("/static", StaticFiles(directory="static"), name="static")

templates = Jinja2Templates(directory="templates")


@app.get("/")
def home(request: Request):
    return templates.TemplateResponse("index.html", {"request":request})
    # return {"message": "hello seoul"}

#
# @app.get("/")
# async def root():
#     return {"message": "Hello World"}
