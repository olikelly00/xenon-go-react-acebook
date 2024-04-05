import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { getPosts, createPosts, updatePostLikes, deletePosts} from "../../services/posts";
import Post from "../../components/Post/Post";
import "./FeedPage.css"


export const FeedPage = () => {
    const [posts, setPosts] = useState([]);
    const [post, setPost] = useState("");
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            getPosts(token)
                .then((data) => {
                    const sortedPosts = data.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
                    setPosts(sortedPosts);
                    localStorage.setItem("token", data.token);
                })
                .catch((err) => {
                    console.error(err);
                    navigate("/login");
                });
    }
    }, [navigate]);

    const handleLike = async (postId) => {
        try {
            await updatePostLikes(token, postId);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setPosts(sortedPosts);
        } catch (err) {
            console.error(err);
        }
    };
      const token = localStorage.getItem("token");
    if (!token) {
      navigate("/login");
      return;
    }

    const handleDelete = async (postId) => {
        try {
            await deletePosts(token, postId);
            const updatedPosts = await getPosts(token);
            const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
            setPosts(sortedPosts);
        } catch (err) {
            console.error(err);
            setErrorMessage("  🤡 nice try bozo! try deleting your own post instead... 😉");
        }
    };
  

    const handleSubmitPost = async (event) => {
      event.preventDefault();
      try {
        await createPosts(token, post);
        const updatedPosts = await getPosts(token);
        const sortedPosts = updatedPosts.posts.sort((a, b) => new Date(b.created_at) - new Date(a.created_at));
        setPosts(sortedPosts);
        setPost("");
        localStorage.setItem("token", updatedPosts.token);
      } catch (err) {
        console.error(err);
      }
    };

    const handlePostChange = (event) => {
      setPost(event.target.value);
    };

    return (
      <div className="feed-container">
        <h2>Posts</h2>
        <form onSubmit={handleSubmitPost}>
          <div className="create-post">
            <input type="text" value={post} onChange={handlePostChange} />
            <input role="submit-button" id="submit" type="submit" value="Submit" />
          </div>
        </form>
        <div className="feed-all-posts" role="feed">
          {posts.map((post) => (
            <div className="feed-post" key={post._id}>
              <Post post={post} token={token} onDelete={handleDelete} onLike={handleLike} user={post.User} />
            </div>
          ))}
        </div>
          {errorMessage && <p className="error-message">{errorMessage}</p>}
      </div>
    );
  };

