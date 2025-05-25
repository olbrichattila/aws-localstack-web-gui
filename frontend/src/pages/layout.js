import { useRef, useEffect } from "react";

const Layout = ({ Component, Icon, title }) => {
    const ref = useRef(null)

    useEffect(() => {
        if (!ref.current) return
        
        ref.current.classList.remove('visible')
        let timeout = setTimeout(() => {
            ref.current.classList.add('visible')
        }, 10)

        return () => {
            clearTimeout(timeout)
        }
       
    }, [title])



    return (
        <>
            <div className="containerHeader">
                <Icon /><span>{title}</span>
            </div>
            <div className="pageContainer" ref={ref}>
                <Component />
            </div>
        </>
    );

}

export default Layout;
