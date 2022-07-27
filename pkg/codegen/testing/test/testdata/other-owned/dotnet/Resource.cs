// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Other.Example
{
    [ExampleResourceType("example::Resource")]
    public partial class Resource : global::Pulumi.CustomResource
    {
        [Output("bar")]
        public Output<string?> Bar { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs? args = null, CustomResourceOptions? options = null)
            : base("example::Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("example::Resource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "example.com/download",
                AdditionalSecretOutputs =
                {
                    "bar",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, options);
        }
    }

    public sealed class ResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("bar")]
        private Input<string>? _bar;
        public Input<string>? Bar
        {
            get => _bar;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bar = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ResourceArgs()
        {
        }
        public static new ResourceArgs Empty => new ResourceArgs();
    }
}
