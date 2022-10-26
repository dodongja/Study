package com.example.demo.service.yMaker;

import com.example.demo.controller.yMaker.request.BoardRequest;
import com.example.demo.entity.yMaker.Board;

import java.util.List;

public interface BoardService {
    public List<Board> list();
    public void save(BoardRequest board);
    public void delete(String[] arrayParam);
    public List<Board> search(BoardRequest board);
    public void modify(BoardRequest boardRequest);
}

