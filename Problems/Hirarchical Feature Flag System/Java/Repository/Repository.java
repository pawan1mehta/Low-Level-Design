package Repository;

import java.util.List;

public interface Repository<T, K> {
    List<T> readAll();
    void add(K id, T channel);
    T get(K id);
    void delete(K id);
}