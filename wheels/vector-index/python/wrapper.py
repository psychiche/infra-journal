class VectorIndex:
    def __init__(self):
        self._ready = False

    def load(self, path: str) -> None:
        # TODO: 绑定 Go 后端或加载本地索引
        self._ready = True

    def search(self, vector, topk: int = 10):
        if not self._ready:
            raise RuntimeError("index not loaded")
        return []
