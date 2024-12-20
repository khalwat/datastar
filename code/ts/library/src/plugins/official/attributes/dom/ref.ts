// Authors: Delaney Gillilan
// Icon: mdi:cursor-pointer
// Slug: Create a reference to an element
// Description: This attribute creates a reference to an element that can be used in other expressions.

import { AttributePlugin } from "../../../../engine";
import { elemToSelector } from "../../../../utils/dom";

// Sets the value of the element
export const RefAttribute: AttributePlugin = {
    pluginType: "attribute",
    name: "ref",
    mustHaveEmptyKey: true,
    mustNotEmptyExpression: true,
    bypassExpressionFunctionCreation: () => true,
    onLoad: (ctx) => {
        ctx.upsertIfMissingFromStore("_dsPlugins.refs", {});
        const { el, expression } = ctx;
        const s = ctx.store();

        const revised = {
            _dsPlugins: {
                refs: {
                    ...s._dsPlugins.refs.value,
                    [expression]: elemToSelector(el),
                },
            },
        };
        ctx.mergeStore(revised);

        return () => {
            const s = ctx.store();
            const revised = { ...s._dsPlugins.refs.value };
            delete revised[expression];
            s._dsPlugins.refs = revised;
        };
    },
};
